package main

import "fmt"

func main() {

	printLine("C for")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	printLine("while")

	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	printLine("infinite loop")

	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	printLine("Skipping")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	printLine("Array")
	languages := [3]string{"Go", "Java", "Python"}
	for index, value := range languages {
		fmt.Println(index, value)
	}

}

func printLine(title string) {
	fmt.Println("--------------", title, "--------------")
}
