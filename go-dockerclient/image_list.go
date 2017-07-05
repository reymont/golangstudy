package main

import (
	"fmt"
	"reflect"
	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
		//fmt.Println(client.InspectContainer("03e33bcbb262714d67fb3f9f19a5c8dc61ab5f35493d81840474d80a72e9fcbb"))
	}
	fmt.Println(client.InspectContainer("f4ea3eb99a56dc358a658ff6fa82324c050a7a5a66c6c20930ceadc4829eaf15"))
	id := "f4ea3eb99a56dc358a658ff6fa82324c050a7a5a66c6c20930ceadc4829eaf15"
	container, _ := client.InspectContainer(id)
	fmt.Println("NetworkMode:---------"+container.HostConfig.NetworkMode)
	s := reflect.Indirect(reflect.ValueOf(container.NetworkSettings))
	networks := s.FieldByName("Networks")
	if networks.IsValid() {
		var ip string
		for _, net := range networks.MapKeys() {
			if net.Interface().(string) == container.HostConfig.NetworkMode {
				ip = networks.MapIndex(net).FieldByName("IPAddress").Interface().(string)
				fmt.Printf("%s %v", net, ip)
			}
			fmt.Println(net)
			fmt.Println(networks.MapIndex(net));
		}

		var networkID string
		for _, net := range networks.MapKeys() {
			if net.Interface().(string) == container.HostConfig.NetworkMode {
				networkID = networks.MapIndex(net).FieldByName("NetworkID").Interface().(string)
				fmt.Printf("%s %v", net, networkID)
			}
		}
	}
	execObj, _ :=client.InspectExec(id)
	fmt.Printf("exec:-------------%v",execObj)
}
