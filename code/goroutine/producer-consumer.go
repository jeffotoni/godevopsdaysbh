package main

import "fmt"

var stream = make(chan int)
var done = make(chan bool)

func produce() {
	for i := 0; i < 10; i++ {
		//fmt.Println("sending")
		stream <- i
		//fmt.Println("sent")
	}

	fmt.Println("No closing channel")
	done <- true
}

func consume() {
	for {
		data := <-stream
		fmt.Println("Consumer: ", data)
	}
}

func main() {

	for j := 0; j < 10000; j++ {
		go produce()
	}

	go consume()

	fmt.Println("After calling DONE")
}
