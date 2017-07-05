package main

import (
	"flag"
	"fmt"
	"os/exec"
	"regexp"
)

//go run docker_inspect_args.go -container="dfdf"
func main() {
	container := flag.String("container", "", "container")
	// flag.Args方式
	flag.Parse()

	fmt.Println("container:", *container)

	f, err := exec.Command("bash", "-c", "docker inspect -f '{{.HostConfig.NetworkMode}}' "+*container).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
	env, _ := exec.Command("bash", "-c", "docker inspect -f '{{.Config.Env}}' "+*container).Output()
	fmt.Println("env", string(env))

	reg1 := regexp.MustCompile(`SERVICE_ID=\w+`)
	service_id_str := reg1.FindAllString(string(env), -1)
	fmt.Println("service_id_str:", service_id_str)

	reg2 := regexp.MustCompile(`DEPLOY_ID=[\w|-]+`)
	deploy_id_str := reg2.FindAllString(string(env), -1)
	fmt.Println("deploy_id_str:", deploy_id_str)
}
