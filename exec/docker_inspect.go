package main

import (
	"fmt"
	"os/exec"
)

func main() {
	f, err := exec.Command("bash", "-c", `docker inspect -f '{{.HostConfig.NetworkMode}}' 3a0a1a1c5e36`).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}
