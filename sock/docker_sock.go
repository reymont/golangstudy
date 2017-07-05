package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
)

type Image struct {
	Created     uint64
	Id          string
	ParentId    string
	RepoTags    []string
	Size        uint64
	VirtualSize uint64
}

func main() {
	addr := net.UnixAddr{"/var/run/docker.sock", "unix"}
	conn, err := net.DialUnix("unix", nil, &addr)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte("GET /images/json HTTP/1.0\r\n\r\n"))
	if err != nil {
		panic(err)
	}

	result, err_conn := ioutil.ReadAll(conn)
	if err_conn != nil {
		panic(err_conn)
	}

	fmt.Println("result:", string(result))

	body := getBody(result[:])

	var images []Image

	err_js := json.Unmarshal(body, &images)
	if err_js != nil {
		panic(err_js)
	}

	fmt.Println("len of images: ", len(images))
	fmt.Println("image.Id: ", images[0].Id)
}
func getBody(result []byte) (body []byte) {
	for i := 0; i <= len(result)-4; i++ {
		if result[i] == 13 && result[i+1] == 10 && result[i+2] == 13 && result[i+3] == 10 {
			body = result[i+4:]
			break
		}
	}
	return
}
