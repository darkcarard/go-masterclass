package main

import (
	"fmt"
	"strings"
)

const (
	division          = "Division"
	divisionByZeroMsg = "division by zero is not allowed"
)

type mathError struct {
	operation string
	inputA    int
	inputB    int
	message   string
}

func (me *mathError) Error() string {
	var inputs []string
	if me.operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", me.inputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", me.inputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s", me.operation,
		strings.Join(inputs, ","), me.message)
}

func main() {
	fmt.Println(sum(1, 2, 3))
	printResult(safeDivision(10, 0))
	printResult(safeDivision(10, 2))
}

func sum(numbers ...int) int {
	defer fmt.Println("Sum finished")

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func safeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &mathError{
			operation: division,
			inputA:    a,
			inputB:    b,
			message:   divisionByZeroMsg,
		}
	}

	return a / b, nil
}

func printResult(value int, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result: ", value)
}
