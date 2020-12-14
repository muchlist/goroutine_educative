package main

import "fmt"

func main() {
	// This code causes deadlock
	//myChannel := make(chan int)
	//myChannel <- 10
	//fmt.Println(<-myChannel)

	//because channel must be written in goroutine
	myChannel := make(chan int)
	go func() {
		myChannel <- 10
	}()
	fmt.Println(<-myChannel)

	//or we can use channel buffering so channel not blocking for define capacity
	myChannel2 := make(chan int, 2)
	myChannel2 <- 10
	myChannel2 <- 20
	// myChannel2 <- 20  //third received channel causes deadlock because cap is just 2
	fmt.Println(<-myChannel2)
}
