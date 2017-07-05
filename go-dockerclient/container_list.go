package main

import (
	"github.com/fsouza/go-dockerclient"
	"fmt"
	"encoding/json"
)
//go run container_list.go |python -mjson.tool
func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	containers, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		fmt.Println("containers is empty")
	}
	b, err := json.Marshal(containers)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	for _,c := range containers {
		fmt.Println(c.ID)
	}
}