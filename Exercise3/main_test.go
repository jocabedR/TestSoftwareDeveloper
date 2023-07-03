package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertEmployee(t *testing.T) {
	insertEmployee(1, "Jocabed", "Ríos", "Saucedo", time.Date(1999, time.May, 16, 0, 0, 0, 0, time.UTC))
	insertEmployee(2, "Ana", "Ríos", "Salcedo", time.Date(2001, time.May, 16, 0, 0, 0, 0, time.UTC))
	insertEmployee(3, "César", "Álvarez", "Fuentes", time.Date(2000, time.May, 16, 0, 0, 0, 0, time.UTC))

	expect1 := Employee{
		ID:        1,
		Name:      "Jocabed",
		LastName1: "Ríos",
		LastName2: "Saucedo",
		BirthDate: time.Date(1999, time.May, 16, 0, 0, 0, 0, time.UTC),
	}
	expect2 := Employee{
		ID:        2,
		Name:      "Ana",
		LastName1: "Ríos",
		LastName2: "Salcedo",
		BirthDate: time.Date(2001, time.May, 16, 0, 0, 0, 0, time.UTC),
	}
	expect3 := Employee{
		ID:        3,
		Name:      "César",
		LastName1: "Álvarez",
		LastName2: "Fuentes",
		BirthDate: time.Date(2000, time.May, 16, 0, 0, 0, 0, time.UTC),
	}

	result1 := Employees[1]
	result2 := Employees[2]
	result3 := Employees[3]

	assert.Equal(t, expect1, result1)
	assert.Equal(t, expect2, result2)
	assert.Equal(t, expect3, result3)
	assert.NotEqual(t, expect1, result3)
}

func TestGetEmployeeAgeById(t *testing.T) {
	insertEmployee(1, "Jocabed", "Ríos", "Saucedo", time.Date(1999, time.May, 16, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, 24, getEmployeeAgeById(1))
}

func TestGetAllOrderByLastName(t *testing.T) {
	expect := []Employee{
		{
			ID:        3,
			Name:      "César",
			LastName1: "Álvarez",
			LastName2: "Fuentes",
			BirthDate: time.Date(2000, time.May, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			Name:      "Ana",
			LastName1: "Ríos",
			LastName2: "Salcedo",
			BirthDate: time.Date(2001, time.May, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        1,
			Name:      "Jocabed",
			LastName1: "Ríos",
			LastName2: "Saucedo",
			BirthDate: time.Date(1999, time.May, 16, 0, 0, 0, 0, time.UTC),
		},
	}

	assert.Equal(t, expect, getAllOrderByLastName())
}

func TestGetAllOrderByAge(t *testing.T) {
	expect := []Employee{
		{
			ID:        2,
			Name:      "Ana",
			LastName1: "Ríos",
			LastName2: "Salcedo",
			BirthDate: time.Date(2001, time.May, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        3,
			Name:      "César",
			LastName1: "Álvarez",
			LastName2: "Fuentes",
			BirthDate: time.Date(2000, time.May, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        1,
			Name:      "Jocabed",
			LastName1: "Ríos",
			LastName2: "Saucedo",
			BirthDate: time.Date(1999, time.May, 16, 0, 0, 0, 0, time.UTC),
		},
	}

	assert.Equal(t, expect, getAllOrderByAge())
}
