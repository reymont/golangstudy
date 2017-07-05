package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"bytes"
	"log"
	"html"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	url_path := html.EscapeString(r.URL.Path[1:])
	if r.Method == "GET" {
		log.Printf("Hello %s", html.EscapeString(r.URL.Path[1:]))
	} else if r.Method == "POST" {
		defer func() { //必须要先声明defer，否则不能捕获到panic异常
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		b, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		log.Printf("POST data : %s\n", b)

		if(strings.HasPrefix(url_path,"config")){
			if(!strings.HasSuffix(url_path,"rules")&&!strings.HasSuffix(url_path,"yml")){
				log.Println(url_path[7:]+" is not config file")
				return
			}
			newConfig(b, url_path[7:])
		}

		if(strings.HasPrefix(url_path,"api")){
			file_bs, file_err := ioutil.ReadFile("/etc/prometheus/alert.targets")
			if file_err != nil {
				log.Println(file_err)
			}
			res,err := http.Post(string(file_bs), "application/json;charset=utf-8", bytes.NewBuffer([]byte(b)))
			//res,err := http.Post("http://192.168.31.178:8080/monitor/api/v1/alerts", "application/json;charset=utf-8", bytes.NewBuffer([]byte(s)))
			if err != nil {
				log.Println(err)
			}
			result, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				log.Println(err)
			}
			w.Write([]byte(result))
			return
		}
		log.Println("Do nothing!")
	}
}

func newConfig(body []byte, f string) {
	if(strings.HasPrefix(f,"/")){
		f = f[1:]
	}

	c, err := os.Create("/etc/prometheus/"+f)
	if err != nil {
		log.Println("writeLines: %s", err)
	}
	defer c.Close()

	log.Println(string(body))
	c.Write(body)
	c.Sync()
	log.Println("rewrite config "+c.Name())
	return
}
