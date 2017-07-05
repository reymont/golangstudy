package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"go-program/go-restful/userservice"
)

func main() {
	restful.Add(userservice.New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}