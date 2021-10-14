/**
存放A/B Test 实验配置和服务配置模块
单机可以存放本地、存放文件服务器、数据库等
分布式环境可以使用分布式数据库，如zoonkeeper、etcd等
这里仅用本地文件作为演示所用
*/
package mydb

import (
	"io/ioutil"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"

	"abtest/pkg/common"
)

var (
	// Whitelist map[project]map[hashkey]map[layerID]zoneID
	Whitelist = make(map[string]map[string]map[string]string)
)

// MustInitServerConfig is a init function
func MustInitServerConfig() *common.ServerConfig {
	b, err := ioutil.ReadFile("../../pkg/mydb/server_template.yaml")
	if err != nil {
		glog.Fatalf("mydb.MustInitConfig: %v", err)
	}
	conf := new(common.ServerConfig)
	if err = yaml.Unmarshal(b, conf); err != nil {
		glog.Fatalf("mydb.MustInitConfig yaml unmarshal: %v", err)
	}
	return conf
}

// MustInitWhitelistConfig  is a init function
func MustInitWhitelistConfig() {
	b, err := ioutil.ReadFile("../../pkg/mydb/whitelist_template.yaml")
	if err != nil {
		glog.Fatalf("mydb.MustInitConfig: %v", err)
	}

	if err = yaml.Unmarshal(b, Whitelist); err != nil {
		glog.Fatalf("mydb.MustInitConfig yaml unmarshal: %v", err)
	}

	glog.Infof("db once sync Whitelist Config: whitelist: %v", Whitelist)
}

// GetABTestConfig is to get abtest lab config
func GetABTestConfig() []byte {
	b, err := ioutil.ReadFile("../../pkg/mydb/zone_template.yaml")
	if err != nil {
		glog.Fatalf("mydb.MustInitConfig: %v", err)
	}

	return b
}
