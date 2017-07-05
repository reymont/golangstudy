package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//执行【ls /】并输出返回文本
	f, err := exec.Command("ls", "/").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}
