package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Before the return of main")
	}()

	simpleDefer()
	lifoSimpleDefer()
	printDottedLine()

	fmt.Println("Last in main")
}

func simpleDefer() {
	printDottedLine()
	fmt.Println("Start")
	defer fmt.Println("Deferred")
	fmt.Println("middle")
}

func lifoSimpleDefer() {
	printDottedLine()
	fmt.Println("Start")
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	fmt.Println("middle")
}

func printDottedLine() {
	fmt.Println("---------------------------------")
}
