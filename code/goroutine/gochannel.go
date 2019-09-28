package main

import (
	"fmt"
	//"time"
)

func main() {

	c1 := make(chan bool, 10)
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	c1 <- true
	// }()

	<-c1
	fmt.Println("done")

}
