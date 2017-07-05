package main

import "fmt"

type count struct {
	total string
	a1 int
	a2 int
}

func main() {
	inter_map := make(map[string]int)
	maps := make(map[string]map[string]int)

	for i:=0 ; i<10 ; i++  {
		inter_map["test1"]++
		inter_map["test"+fmt.Sprintf("%d", 2) ]++
		maps["test"] = inter_map
	}

	fmt.Println(maps)
	fmt.Println(len(maps))
	fmt.Println(string(2))
}
