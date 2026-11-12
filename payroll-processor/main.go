package main

import "fmt"

type payable interface {
	fmt.Stringer
	calculatePay() float64
}

type salaryEmployee struct {
	name         string
	annualsalary float64
}

func (se salaryEmployee) calculatePay() float64 {
	return se.annualsalary / 12
}

func (se salaryEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)",
		se.name, se.annualsalary)
}

type hourlyEmployee struct {
	name        string
	hourlyRate  float64
	hoursWorked float64
}

func (he hourlyEmployee) calculatePay() float64 {
	return he.hourlyRate * he.hoursWorked
}

func (he hourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f, Hours: %.1f)",
		he.name, he.hourlyRate, he.hoursWorked)
}

type comissionEmployee struct {
	name          string
	baseSalary    float64
	comissionRate float64
	salesAmmount  float64
}

func (ce comissionEmployee) calculatePay() float64 {
	return ce.baseSalary + (ce.comissionRate * ce.salesAmmount)
}

func (ce comissionEmployee) String() string {
	return fmt.Sprintf("Comission: %s (Base: %.2f, CommRate: $%.2f, Sales: %.1f)",
		ce.name, ce.baseSalary, ce.comissionRate, ce.salesAmmount)
}

func printEmployeeSumary[s fmt.Stringer](employee s) {
	fmt.Printf(" - Processing: %s\n", employee)
}

func processPayroll(employees []payable) {
	fmt.Println("\n--- Processing Payroll ---")
	totalPayroll := 0.0
	for _, employee := range employees {
		printEmployeeSumary(employee)
		pay := employee.calculatePay()
		fmt.Printf("   Monthly Pay: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nTotal Monthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("--------------------------------")
}

func main() {
	fmt.Println("Welcome to Payroll Processor!")
	salEmp := salaryEmployee{name: "Alice", annualsalary: 72000.00}
	hrEmp := hourlyEmployee{name: "Bob", hourlyRate: 25.00, hoursWorked: 160.0}
	comEmp := comissionEmployee{name: "Charlie", baseSalary: 2000.00,
		comissionRate: 0.10, salesAmmount: 15000.00}

	payrollList := []payable{
		salEmp,
		hrEmp,
		comEmp,
		hourlyEmployee{name: "Diana", hourlyRate: 30.00, hoursWorked: 150.0},
	}

	processPayroll(payrollList)
}
