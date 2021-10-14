package server

import (
	"context"
	"fmt"

	"github.com/golang/glog"

	pb "abtest/abtest_proto"
	db "abtest/pkg/cache"
	"abtest/pkg/hash"
	"abtest/pkg/mydb"
	"abtest/pkg/prom"
)

// GetABTestZone is to get the ABTest target zone
func (s *ABTestGRPCServer) GetABTestZone(ctx context.Context, r *pb.GetABTestZoneRequest) (*pb.GetABTestZoneResponse, error) {
	response := &pb.GetABTestZoneResponse{}

	if r.GetProjectId() == "" {
		return response, fmt.Errorf("Param ProjectId empty")
	}

	if r.GetLayerId() == "" {
		return response, fmt.Errorf("Param LayerId empty")
	}

	if r.GetHashKey() == "" {
		return response, fmt.Errorf("Param HashKey empty")
	}

	response.Zone = matchZone(r.GetProjectId(), r.GetLayerId(), r.GetHashKey())

	// add data into prom counter
	prom.AddPromCounter(r.GetProjectId(), r.GetLayerId(), response.Zone.GetId(), response.Zone.GetLabel())

	return response, nil
}

func matchZone(projectID, layerID, hashkey string) *pb.Zone {

	// get zones from cache
	zones := db.GetZones(projectID, layerID)
	// glog.Info("matchZone zones:", zones)

	if len(zones) == 0 {
		return nil
	}

	// check whitelist
	if ok, zone := whitlist(projectID, layerID, hashkey, zones); ok {
		glog.Infof("Project(%s) layerID(%s) HashKey(%s) Whitelist bingo! TargetZone: %+v", projectID, layerID, hashkey, zone)
		return zone
	}

	// hash to match the target zone
	hashValue := hash.Hash(hashkey, layerID, uint32(zones[0].GetLayer().GetTotalWeight()))

	// check current zone
	for _, zone := range zones {
		if zone.GetWeight().GetMax() >= int32(hashValue) && zone.GetWeight().GetMin() <= int32(hashValue) {
			// check the parent zone and make sure the user comes from parent zones
			isFromParent := false
			for _, parent := range zone.GetLayer().GetParentZones() {
				if parent.GetId() == matchZone(parent.GetProject().GetId(), parent.GetLayer().GetId(), hashkey).GetId() {
					isFromParent = true
					break
				}
			}

			// return the matched zone
			if len(zone.GetLayer().GetParentZones()) == 0 || isFromParent {
				return zone
			}
		}
	}
	return nil
}

func whitlist(projectID, layerID, hashkey string, zones []*pb.Zone) (bool, *pb.Zone) {
	if project2Whitelist, ok := mydb.Whitelist[projectID]; ok {
		for targetLayerId, targetZoneID := range project2Whitelist[hashkey] {
			if layerID == targetLayerId {
				for _, zone := range zones {
					if zone.GetId() == targetZoneID {
						return true, zone
					}
				}
			}
		}
	}

	return false, nil
}
