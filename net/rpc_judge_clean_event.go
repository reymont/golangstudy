package main

import (
	"fmt"

	"net"

	cmodel "github.com/open-falcon/common/model"
	"net/rpc"
	"log"
	"time"
)

func main() {
	addr := "192.168.0.179:6080"
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

	resp := &cmodel.SimpleRpcResponse{}
	err = rpcClient.Call("Judge.CleanEvent", "e_487_1e260b72d9b46fd357378893e7f8616a", &resp)

	if err != nil {
		if rpcClient != nil {
			rpcClient.Close()
			rpcClient = nil
		}
	}
}
