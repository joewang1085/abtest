package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	db "abtest/pkg/cache"
	"abtest/pkg/mydb"
	"abtest/pkg/mymq"
	"abtest/pkg/prom"
	"abtest/pkg/server"

	"github.com/golang/glog"
)

func main() {

	_ = flag.Set("logtostderr", "true")
	flag.Parse()

	// global ctx
	gctx, cancer := context.WithCancel(context.Background())
	defer cancer()

	// get config from dva
	conf := mydb.MustInitServerConfig()

	// get whitelist from dva
	mydb.MustInitWhitelistConfig()

	// set duration to sync db
	glog.Info("cache is syncing db ...")
	duration := time.Second * 60
	if conf.DBConfig.Duration != 0 {
		duration = time.Second * time.Duration(conf.DBConfig.Duration)
	}
	db.SyncABTestConfigDBTask(gctx, duration) // sync ABTest/Whitelist config

	// init mq producer
	mymq.Produce(gctx)

	// start prometheus server
	glog.Info("Prometheus server is listening ...")
	if conf.PrometheusConfig.Address != "" {
		prom.PromAddress = fmt.Sprintf(":%s", conf.PrometheusConfig.Address)
	}
	go prom.StartPromServer()

	// start grpc server
	glog.Info("GRPC server is listening ...")
	if conf.GRPCConfig.Address != "" {
		server.GRPCAddress = fmt.Sprintf(":%s", conf.GRPCConfig.Address)
	}
	server.MustStartGRPCServer()
}
