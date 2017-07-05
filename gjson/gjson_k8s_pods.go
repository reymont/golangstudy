package main

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	URL := "http://192.168.0.180:8080/api/v1/namespaces/716mo7m10myxdnfbgfpr3f3dh8npzk/pods/apidemorest4lp2ggpc-vqg5j"

	var result []byte
	if response, err := http.Get(URL); err != nil {
		println(err)
	} else {
		if result, err = ioutil.ReadAll(response.Body); err != nil {
			println(err)
		}
	}

	//println(string(result))
	value := gjson.Get(string(result), "status.phase")
	println(value.String())
	println(gjson.Get(string(result), "status.hostIP").String())
}
