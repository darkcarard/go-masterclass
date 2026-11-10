package main

import "fmt"

func main() {
	age := 10
	fmt.Printf("Age: %p\n", &age)
	fmt.Printf("Age: %d\n", &age) //Decimal format

	num := 10
	fmt.Printf("Value: %+v\n", num)
	modifyValue(num)
	fmt.Printf("Value: %+v\n", num)
	modifyValueByPointer(&num)
	fmt.Printf("Value: %+v\n", num)
}

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("Value modified: %+v\n", val)
}

func modifyValueByPointer(val *int) {
	if val == nil {
		fmt.Println("VAL is nil")
	}

	*val = *val * 10
	fmt.Printf("Value modified: %+v\n", *val)
}
