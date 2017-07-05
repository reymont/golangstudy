package main

import (
	"fmt"
	"net"
	"log"
	"time"
	"net/rpc"

	"github.com/open-falcon/common/model"
)

func main()  {
	//addr := "192.168.0.179:6030"
	addr := "192.168.31.56:6030"
	_, err := net.ResolveTCPAddr("tcp",addr)
	if err != nil {
		log.Println(addr, "format error", err)
	}

	conn, err := net.DialTimeout("tcp", addr,  2*time.Second)
	if err != nil {
		log.Printf("new conn fail, addr %s, err %v", addr, err)
		return
	}

	rpcClient := rpc.NewClient(conn)

	if err != nil {
		fmt.Errorf("ConnectError: %s", err.Error())
	}
	defer rpcClient.Close()


	req := model.AgentReportRequest{
		Hostname:      "192.168.31.39",
		IP:            "192.168.31.31",
		AgentVersion:  "5.1.1",
		PluginVersion: "plugin not enabled",
		CpuNum:        "2",
		MemNum:        "3776",
	}

	var resp model.SimpleRpcResponse
	err = rpcClient.Call("Agent.ReportStatus", req, &resp)
	if err != nil || resp.Code != 0 {
		log.Println("call Agent.ReportStatus fail:", err, "Request:", req, "Response:", resp)
	}

	if err != nil {
		if rpcClient != nil {
			rpcClient.Close()
			rpcClient = nil
		}
	}
}
