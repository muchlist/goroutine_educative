package main

import (
	"fmt"
	"sync"
	"time"
)

func printTable(n int, wg *sync.WaitGroup) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for number := 2; number <= 12; number++ {
		wg.Add(1)
		go printTable(number, &wg)
	}

	wg.Wait()
	/*
		1 x 12 = 12
		1 x 7 = 7
		1 x 8 = 8
		1 x 9 = 9
		1 x 10 = 10
		1 x 11 = 11
		1 x 6 = 6
		1 x 5 = 5
		1 x 2 = 2
		1 x 4 = 4
		1 x 3 = 3
		2 x 4 = 8
		2 x 3 = 6
		2 x 7 = 14
		2 x 12 = 24
		2 x 8 = 16
		2 x 10 = 20
		2 x 11 = 22
		2 x 6 = 12
		2 x 2 = 4
		...
	*/
}
