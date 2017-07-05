package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	cmd := exec.Command("bash", "-c", `ps -ef | grep -v "grep" | grep "ping"`)
	cmd.Start()
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("Output: " + err.Error())
		return
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	fmt.Printf("stdout: %s", bytes)

}

// 运行命令: go run main.go
