package main

import (
	"flag"

	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"

	//"github.com/golang/glog"
	"fmt"
)

func staticClientExample() {
	staticClient, err := client.NewClient("http://192.168.0.186:4194")
	if err != nil {
		//glog.Errorf("tried to make client and got error %v", err)
		return
	}
	//einfo, err := staticClient.EventStaticInfo("?oom_events=true")
	machineInfo, err := staticClient.MachineInfo()
	if err != nil {
		fmt.Printf("got error retrieving event info: %v", err)
		return
	}
	fmt.Println(machineInfo)
	dockerContainer, err := staticClient.DockerContainer("/be5792e17e67a61c40b810e3baf74227a414346588e89d035d3918ffe8f17557",&info.ContainerInfoRequest{
		NumStats: 1,
	})
	if err != nil {
		fmt.Printf("got error retrieving event info: %v", err)
		return
	}
	fmt.Println(dockerContainer)
	containerInfo, err := staticClient.ContainerInfo("/docker",&info.ContainerInfoRequest{
		NumStats: 1,
	})
	if err != nil {
		fmt.Printf("got error retrieving event info: %v", err)
		return
	}
	fmt.Println(containerInfo)
	//for idx, ev := range einfo {
	//	fmt.Printf("static einfo %v: %v", idx, ev)
	//}
}

func streamingClientExample(url string) {
	streamingClient, err := client.NewClient("http://192.168.0.186:4194")
	if err != nil {
		fmt.Errorf("tried to make client and got error %v", err)
		return
	}
	einfo := make(chan *info.Event)
	go func() {
		err = streamingClient.EventStreamingInfo(url, einfo)
		if err != nil {
			fmt.Errorf("got error retrieving event info: %v", err)
			return
		}
	}()
	for ev := range einfo {
		fmt.Printf("streaming einfo: %v\n", ev)
	}
}

// demonstrates how to use event clients
func main() {
	flag.Parse()
	staticClientExample()
	streamingClientExample("?creation_events=true&stream=true&oom_events=true&deletion_events=true")
}
