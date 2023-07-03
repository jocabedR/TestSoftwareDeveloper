package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	LastName1 string
	LastName2 string
	BirthDate time.Time
}

func newEmployee(id int, name, lastName1, lastName2 string, birthDate time.Time) Employee {
	return Employee{
		ID:        id,
		Name:      name,
		LastName1: lastName1,
		LastName2: lastName2,
		BirthDate: birthDate,
	}
}

func getEmployeeById(id int) Employee {
	employee, exists := Employees[id]

	if !exists {
		fmt.Println("Employee not founded!")
		menu()
	}

	return employee
}
