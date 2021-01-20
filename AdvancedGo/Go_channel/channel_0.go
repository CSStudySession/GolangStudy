package main

import "fmt"

func main() {
	ch0 := make(chan int, 3)
	ch0 <- 2
	ch0 <- 1
	ch0 <- 3
	elem0 := <-ch0
	fmt.Printf("The first element received from channel ch0: %v\n", elem0)

	ch1 := make(chan int, 1)
	ch1 <- 1
	// ch1 <- 2 // full channel, blocking in here.

	ch2 := make(chan int, 1)
	// elem, ok := <-ch2 // empty channel, blocking in here.
	// _, _ = elem, ok
	ch2 <- 1

	var ch3 chan int
	// ch3 <- 1 // nil value, blocking forever!
	// <-ch3 // nil value, blocking forever!
	_ = ch3
}
