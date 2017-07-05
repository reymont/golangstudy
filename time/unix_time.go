package main

import (
	"time"
	"fmt"
)

func main(){
	nowTs := time.Now().Unix()
	lastUpTs := nowTs - nowTs%int64(60)
	rra1StartTs := lastUpTs - int64(720*60)
	fmt.Println(nowTs)
	fmt.Println(nowTs%int64(60))
	fmt.Println(lastUpTs)
	fmt.Println(rra1StartTs)
}
