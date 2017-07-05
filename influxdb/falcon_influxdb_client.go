package main

import (
	cpool "github.com/open-falcon/transfer/sender/conn_pool"
	"github.com/siddontang/go/log"
	"github.com/open-falcon/common/model"
)

var (
	JudgeConnPools     *cpool.SafeRpcConnPools
	TsdbConnPoolHelper *cpool.TsdbConnPoolHelper
	GraphConnPools     *cpool.SafeRpcConnPools
	InfluxDBConnPoolHelper *cpool.InfluxDBConnPoolHelper
)

func main() {
	InfluxDBConnPools := cpool.CreateSafeRpcConnPools(32, 32,
		1000, 5000, []string{"127.0.0.1:6080"})
	log.Debug(InfluxDBConnPools)
	log.Debug(InfluxDBConnPools.M)

	InfluxDBConnPoolHelper = cpool.NewInfluxDBConnPoolHelper("http://192.168.0.179:8086", 32, 32, 1000, 5000)
	m := &model.MetricValue{
		Endpoint:"192.168.31.171",
		Metric:"cpu_usage.cpu_usage",
		//Tags:map[string]string{"cpu": "cpu-total"},
		Value:map[string]interface{}{
			"idle":   10.1,
			"system": 53.3,
			"user":   46.6,
			}}
	InfluxDBConnPoolHelper.Send(m)

}

/**
type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}
 */

/**
 "graph": {
        "enabled": true,
        "batch": 200,
        "connTimeout": 1000,
        "callTimeout": 5000,
        "maxConns": 32,
        "maxIdle": 32,
        "replicas": 500,
        "cluster": {
            "graph-00" : "127.0.0.1:6070"
        }
    },
 */
