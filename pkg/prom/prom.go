package prom

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	diversion = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Diversion", // TODO : 增加 namespace 等方便抓数据
			Help: "Diversion Count",
		},
		[]string{"ProjectId", "LayerId", "ZoneLabel"},
	)

	// PromAddress ...
	PromAddress string = ":8080"
)

// AddPromCounter is to increase prometheus counter
func AddPromCounter(projectID, layerID, zoneID, ZoneLabel string) {
	diversion.With(prometheus.Labels{"ProjectId": projectID, "LayerId": layerID, "ZoneLabel": ZoneLabel}).Inc()
}

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(diversion)
}

// StartPromServer is to start prometheus server
func StartPromServer() {

	http.Handle("/metrics", promhttp.Handler())
	glog.Fatal(http.ListenAndServe(PromAddress, nil))
}
