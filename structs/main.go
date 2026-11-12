package main

import (
	"fmt"
	"go-masterclass/structs/models"
)

func main() {
	pepito := models.NewEmployee(1, "Pepito",
		"Pérez", "Developer", 2000, true)
	fmt.Printf("%+v\n", pepito)

	sara := models.NewEmployee(1, "Sara",
		"Perez", "Manager", 5000, true)

	fmt.Printf("%+v\n", sara)
	fmt.Println(sara.FullName())
	sara.Deactivate()
	fmt.Println(sara)
}
