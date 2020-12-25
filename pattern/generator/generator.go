package main

import "fmt"

func updatePosition(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name, i)
		}
	}()

	return positionChannel
}

func fibonacci(n int) chan int {
	myChannel := make(chan int)
	go func() {
		k := 0
		for i, j := 0, 1; k < n; k++ {
			myChannel <- i
			i, j = i+j, i

		}
		close(myChannel)
	}()
	return myChannel
}

func main() {

	//GENERATOR 1
	positionChannel1 := updatePosition("Legolas :")
	positionChannel2 := updatePosition("Gandalf :")

	for i := 0; i < 5; i++ {
		fmt.Println(<-positionChannel1)
		fmt.Println(<-positionChannel2)
	}

	fmt.Println("Done with getting updates on positions.")

	/*
		Legolas : 0
		Gandalf : 0
		Legolas : 1
		Gandalf : 1
		Legolas : 2
		Gandalf : 2
		Legolas : 3
		Gandalf : 3
		Legolas : 4
		Gandalf : 4
		Done with getting updates on positions.
	*/

	//GENERATOR 2 FIBONACI
	for i := range fibonacci(10) {
		//do anything with the nth term while the fibonacci()
		//is computing the next term
		fmt.Println(i)
	}

	/*
		0
		1
		1
		2
		3
		5
		8
		13
		21
	*/
}
