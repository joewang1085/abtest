package db

import (
	"context"
	"sync"
	"time"

	pb "abtest/abtest_proto"
	"abtest/pkg/mydb"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

var (
	// ZonesCache caches the zones
	ZonesCache sync.Map
)

// GetZones is to describe zones from the cached zones by projectID and layerID
func GetZones(projectID, layerID string) []*pb.Zone {
	zones := make([]*pb.Zone, 0)
	if projectID == "" || layerID == "" {
		return zones
	}

	if projectID2Zones, ok := ZonesCache.Load(projectID); ok {
		// return zones of the layer
		for _, zone := range projectID2Zones.([]*pb.Zone) {
			if layerID == zone.GetLayer().GetId() {
				zones = append(zones, zone)
			}
		}
	}

	return zones
}

// GetOriginLayer is to describe the origin layer from the cache by projectID
func GetOriginLayer(projectID string) *pb.Layer {
	if projectID == "" {
		return nil
	}

	if projectID2Zones, ok := ZonesCache.Load(projectID); ok {
		// return the origin layer
		for _, zone := range projectID2Zones.([]*pb.Zone) {
			if zone.GetLayer().GetParentZones() == nil {
				return zone.GetLayer()
			}
		}
	}

	return nil
}

// GetNextLayer is to describe the next layer from the cached zones by projectID and zoneID
func GetNextLayer(projectID, zoneID string) *pb.Layer {
	if projectID == "" || zoneID == "" {
		return nil
	}

	if projectID2Zones, ok := ZonesCache.Load(projectID); ok {
		// return the next layer
		for _, zone := range projectID2Zones.([]*pb.Zone) {
			if isContained(zone.GetLayer().GetParentZones(), zoneID) {
				return zone.GetLayer()
			}
		}
	}

	return nil
}

func isContained(zones []*pb.Zone, zoneID string) bool {
	for _, zone := range zones {
		if zone.GetId() == zoneID {
			return true
		}
	}

	return false
}

// SyncABTestConfigDBTask is a Synchronizing ABTest config task
func SyncABTestConfigDBTask(ctx context.Context, duration time.Duration) {
	doSyncABTestConfig()

	go func() {
		for true {
			select {
			case <-ctx.Done():
				glog.Info("SyncABTestConfigDBTask done and return")
				return
			case <-time.Tick(duration):
				doSyncABTestConfig()
				mydb.MustInitWhitelistConfig() // whitelist 支持热更新: 定时同步 whitelist
				glog.Info("SyncABTestConfigDBTask once task completed")
			}
		}
	}()
}

func doSyncABTestConfig() {
	config := make(map[string][]*pb.Zone, 0)
	b := mydb.GetABTestConfig()
	if err := yaml.Unmarshal(b, config); err != nil {
		glog.Fatalf("call yaml unmarshal: %v", err)
	}

	for k, v := range config {
		ZonesCache.Store(k, v)
		for i, zone := range v {
			glog.Infof("project:%s,num:%d,zone:%+v", k, i, zone)
		}
	}
}
