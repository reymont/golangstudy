package main

import (
	"strings"
	"strconv"
	"fmt"
)

func main(){
	str := "avg(#1)"
	idx := strings.Index(str, "#")
	limit, _ := strconv.ParseInt(str[idx+1:len(str)-1], 10, 64)
	fmt.Println(idx)
	fmt.Println(str[idx+1:len(str)-1])
	fmt.Println(limit)
}
