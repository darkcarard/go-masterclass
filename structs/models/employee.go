package models

type Employee struct {
	id        int
	firstName string
	lastName  string
	position  string
	salary    int
	isActive  bool
}

func NewEmployee(id int, firstName, lastName, position string, salary int, isActive bool) Employee {
	return Employee{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		position:  position,
		salary:    salary,
		isActive:  isActive,
	}
}

func (e Employee) FullName() string {
	return e.firstName + " " + e.lastName
}

func (e *Employee) Deactivate() {
	e.isActive = false
}
