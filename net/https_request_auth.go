package main

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"os"
	"io/ioutil"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//response, err := client.Get("https://admin:123456@192.168.31.221:6443/")
	response, err := client.Get("https://admin:123456@192.168.31.221:6443/api/v1/nodes")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}