package main

import "fmt"

type person interface {
	getName() string
}

type employee struct {
	id   int
	name string
}

type businessPerson struct {
	id   int
	name string
}

func (e employee) getName() string {
	return e.name
}

func (p employee) String() string {
	return "Employee Name: " + p.getName()

}

func (bp businessPerson) getName() string {
	return bp.name
}

func displayPerson(p person) {
	fmt.Println(p)
}

func main() {
	joe := employee{
		id:   1,
		name: "Joe",
	}

	jane := businessPerson{
		id:   1,
		name: "Jane",
	}

	displayPerson(joe)
	fmt.Println(joe)
	displayPerson(jane)
	fmt.Println(jane)
}
