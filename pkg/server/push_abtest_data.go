package server

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	pb "abtest/abtest_proto"
	db "abtest/pkg/cache"
	"abtest/pkg/mymq"
)

// PushABTestData is to push abtest data to mq
func (s *ABTestGRPCServer) PushABTestData(ctx context.Context, r *pb.PushABTestDataRequest) (*pb.PushABTestDataResponse, error) {
	response := &pb.PushABTestDataResponse{}

	if r.GetProjectId() == "" {
		return response, fmt.Errorf("param ProjectId empty")
	}

	if r.GetHashKey() == "" {
		return response, fmt.Errorf("param HashKey empty")
	}

	zonePath := make([]*pb.Zone, 0)

	layer := db.GetOriginLayer(r.GetProjectId())
	for layer != nil {
		targetZone := matchZone(r.GetProjectId(), layer.GetId(), r.GetHashKey())
		zonePath = append(zonePath, targetZone)
		layer = db.GetNextLayer(r.GetProjectId(), targetZone.GetId())
	}

	// push to mq. 顺序非阻塞发送.
	pushToMQ(r.GetProjectId(), r.GetHashKey(), r.GetKeyType(), r.GetLoginId(), zonePath, r.GetData())

	return response, nil
}

func pushToMQ(projectID, key, keyType, loginID string, zonePath []*pb.Zone, data []*pb.LabData) {

	// strategy = A;B;C
	strategy := ""
	for _, zone := range zonePath {
		if strategy == "" {
			strategy = zone.GetLabel()
		} else {
			strategy = fmt.Sprintf("%s;%s", strategy, zone.GetLabel())
		}
	}

	// extData = k1:v1;k2:v2;k3:v3
	extData := ""
	for _, one := range data {
		if extData == "" {
			extData = fmt.Sprintf("%s:%s", one.Key, one.Value)
		} else {
			extData = fmt.Sprintf("%s;%s", extData, fmt.Sprintf("%s:%s", one.Key, one.Value))
		}
	}

	// send message to mq
	message := pb.ABTestMessage{
		ProjectId:   projectID,
		HashKey:     key,
		KeyType:     keyType,
		Strategy:    strategy,
		Ext:         extData,
		LoginId:     loginID,
		CreatedTime: time.Now().String(),
	}

	msg, err := proto.Marshal(&message)
	if err != nil {
		glog.Errorf("mq message proto to json failed: %v", err)
	} else {
		mymq.SendMessage(fmt.Sprintf("%s:%s", message.GetProjectId(), message.GetHashKey()), string(msg))
	}
}
