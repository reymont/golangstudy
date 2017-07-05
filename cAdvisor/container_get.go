package main

import (
	"io/ioutil"
	"net/http"
	info "github.com/google/cadvisor/info/v1"

	"github.com/tidwall/gjson"
	"client-go/1.4/pkg/util/json"
	"fmt"
)

func main() {
	//URL := "http://192.168.0.180:8080/api/v1/namespaces/716mo7m10myxdnfbgfpr3f3dh8npzk/pods/apidemorest4lp2ggpc-vqg5j"
	URL := "http://192.168.0.181:4194/api/v1.0/containers/system.slice/docker-03cc2d2b5c8aa962a1a04cd8a3d6a81af9e1bda641b58d6b315f4c7ce7658d5d.scope"

	var result []byte
	response, err := http.Get(URL)
	if  err != nil {
		println(err)
	} else {
		if result, err = ioutil.ReadAll(response.Body); err != nil {
			println(err)
		}
	}

	fmt.Println(string(result))

	var s info.ContainerInfo
	json.Unmarshal(result,&s)
	fmt.Println(&s)


	//println(string(result))
	value := gjson.Get(string(result), "status.phase")
	fmt.Println(value.String())
	fmt.Println(gjson.Get(string(result), "status.hostIP").String())
}
