package main

import (
	"fmt"
)

func main() {
	mightPanic(false)
	recoverable()
}

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong!!!!")
	}

	fmt.Println("Executed without panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()

	mightPanic(true)
}
