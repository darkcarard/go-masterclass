package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1.23, 2.87, 3.98))
	fmt.Println(multiply(1, 2, 3))
	fmt.Println(multiply(1.5, 2))
}

func sum[t int | float32 | float64](numbers ...t) t {
	var total t
	for _, number := range numbers {
		total += number
	}

	return total
}

type number interface {
	int | float32 | float64
}

func multiply[t number](numbers ...t) t {
	var total t
	total = 1
	for _, number := range numbers {
		total *= number
	}

	return total
}
