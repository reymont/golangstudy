package main

import "fmt"

func main(){
	map_new := make(map[string]int)
	map_new["ddddd"]++
	map_new["ddddd"]++
	map_new["ddddd"]++
	map_new["aaaaa"]++
	map_new["bbbbb"]++
	fmt.Println(map_new)
}