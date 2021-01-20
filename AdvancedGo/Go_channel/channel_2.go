package main

import "fmt"

var channels = [3]chan int {
	nil,
	make(chan int),
	nil,
}

var nums = []int {1,2,3}

func main() {
	select {
		case getChan(0) <- getNum(0):
			fmt.Printf("the first case is selected.")
		case getChan(1) <- getNum(1):
			fmt.Printf("the second case is selected.")
		case getChan(2) <- getNum(2):
			fmt.Printf("the third case is selected.")
		default:
			fmt.Println("default case is selected.")
	}
}

func getNum(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return nums[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
