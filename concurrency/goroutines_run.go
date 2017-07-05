package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}



	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

	time.Sleep(1*time.Second)

	for _, v := range values {
		fmt.Println(v)
	}
}
