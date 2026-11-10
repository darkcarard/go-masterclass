package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Printf("Grades: %+v\n", studentGrades)

	studentGrades["alice"] = 95
	fmt.Printf("Grades: %+v\n", studentGrades)

	alice, ok := studentGrades["alice"]
	if ok {
		fmt.Printf("Alice: %+v\n", alice)
	}
}
