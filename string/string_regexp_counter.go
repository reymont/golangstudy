package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	counter := "192.168.0.187/container.mem.usage.percent deploy_id=71h668c81k3kpio9qctbtjtq2uphlkz,id=85155ba1c6f9f25d7f5703ac72fee823c1175a5cba342056c2c1b02b3b886339"
	fmt.Printf("%q\n", getDeployIdValue(counter))

	counter2 := "192.168.0.187/container.mem.usage.percent "
	fmt.Printf("%q\n", getDeployIdValue(counter2))
}

func getDeployIdValue(counter string) string {
	reg := regexp.MustCompile(`deploy_id=\w+`)
	str := reg.FindAllString(counter, -1)
	if str == nil {
		return ""
	}
	return strings.Split(str[0], "=")[1]
}
