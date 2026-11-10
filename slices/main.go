package main

import "fmt"

func main() {
	names := []string{"Alice", "John", "Mark"}
	fmt.Printf("Names: %+v\n", names)

	items := make([]int, 3, 5)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)
	items = append(items, 5)
	items = append(items, 6)
	items = append(items, 7)
	items = append(items, 8)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
}
