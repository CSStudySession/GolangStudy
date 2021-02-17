package main

import "fmt"

func main() {
	myPanic()
	defer func(){
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	} ()
	fmt.Println("main method is executed completely.")
}

func myPanic() {
	panic("This is panic!")
	fmt.Println("The function finished execution.")
}