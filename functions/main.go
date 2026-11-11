package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(factorial(5))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(divide(2, 0))
	fmt.Println(divide(10, 2))
	firstName, lastName := splitFullName("Carlos Ardila")
	fmt.Println(firstName)
	fmt.Println(lastName)

}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}

	return a / b, nil
}

func splitFullName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return
}
