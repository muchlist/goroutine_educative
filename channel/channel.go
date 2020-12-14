package main

import "fmt"

func sendValuesWithCloseInside(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	close(myIntChannel)
}

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}

}

func main() {
	myIntChannel := make(chan int)

	go sendValuesWithCloseInside(myIntChannel)

	for i := 0; i < 6; i++ { //loop until 6 for testing closing channel successful
		value, open := <-myIntChannel
		if !open {
			break
		}
		fmt.Println(value) //receiving value
		/*
			0
			1
			2
			3
			4
		*/
	}

	myIntChannel2 := make(chan int)
	// good practice to defer the closing of channels in the main program so that we clean up everything ourselves.
	defer close(myIntChannel2)
	go sendValues(myIntChannel2)

	for i := 0; i < 5; i++ {
		fmt.Println(<-myIntChannel2) //receiving value
		/*
			0
			1
			2
			3
			4
		*/
	}
}
