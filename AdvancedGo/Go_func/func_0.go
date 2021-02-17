package main

import (
	"errors"
	"fmt"
)

type calculatorFunc func(x int, y int) (int, error)
type operate func(x, y int) int

func main() {
	x, y := 56, 32
	op := func(x,y int) int {
		return x + y
	}

	add := genCalculator(op)
	ret, err := add(x, y)
	fmt.Printf("The ret: %d, (err: %v)\n", ret, err)
}

func genCalculator(op operate) calculatorFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}