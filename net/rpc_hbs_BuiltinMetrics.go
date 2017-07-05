package main

import (
	"fmt"
	"log"
	"time"
	"github.com/toolkits/net"
	"github.com/open-falcon/common/model"
)

func main() {
	addr := "192.168.0.179:6030"
	//addr := "192.168.31.56:6030"

	rpcClient, err := net.JsonRpcClient("tcp", addr, 2 * time.Second)

	if err != nil {
		fmt.Errorf("ConnectError: %s", err.Error())
	}
	defer rpcClient.Close()

	var checksum string = ""

	req := model.AgentHeartbeatRequest{
		Hostname:      "192.168.0.179",
		Checksum:       checksum,
	}

	var resp model.BuiltinMetricResponse
	err = rpcClient.Call("Agent.BuiltinMetrics", req, &resp)
	if err != nil {
		log.Println("call Agent.BuiltinMetrics fail:", err, "Request:", req, "Response:", resp)
	}
	log.Println("Response:", resp)


	if err != nil {
		if rpcClient != nil {
			rpcClient.Close()
			rpcClient = nil
		}
	}
}
