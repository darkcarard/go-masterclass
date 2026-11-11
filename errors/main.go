package main

import (
	"errors"
	"fmt"
)

var errDivisionByZero = errors.New("division by zero")
var errNumTooLarge = errors.New("number too large")

func main() {
	printResult(divide(1001, 1))
	printResult(divide(500, 0))
	printResult(divide(500, 5))

}

func printResult(value int, err error) {
	if err != nil {
		if errors.Is(err, errDivisionByZero) {
			fmt.Println("divide by zero")
		} else if errors.Is(err, errNumTooLarge) {
			fmt.Println("number too large")
		}

		return
	}

	fmt.Println(value)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errDivisionByZero
	}

	if a > 1000 {
		return 0, errNumTooLarge
	}

	return a / b, nil
}
