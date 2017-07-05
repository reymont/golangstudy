package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"encoding/json"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	info, err := client.Info()
	if err != nil {
		fmt.Println("containers is empty")
	}
	b, err := json.Marshal(info)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
