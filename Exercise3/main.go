package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
	"unicode"

	age "github.com/theTardigrade/golang-age"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// I decided to use a Map to in the future can look easilly for ID existence.
var Employees = make(map[int]Employee)

func main() {
	menu()
}

func menu() {
	var option int
	fmt.Println("\nPlease, choose an option.")
	fmt.Println("(1) Insert an employee." +
		"\n(2) Get Employee's age by ID" +
		"\n(3) Get Employees list sorted alphabetically by last name." +
		"\n(4) Get Employees list sorted by age." +
		"\n(5) Log out.")

	fmt.Scanln(&option)

	switch option {
	case 1:
		var id int
		var strId, name, lastName1, lastName2, strBirthDate string
		var birthDate time.Time
		var err error

		fmt.Print("\nPlease insert Employee's id: ")
		fmt.Scanln(&strId)
		id, err = strconv.Atoi(strId)
		// If an ID enter cannot be convert into a integer returns to the menu.
		if err != nil {
			fmt.Println("Something was wrong, plis check your enter.")
			menu()
		}
		_, exists := Employees[id]
		// To check if the ID is already used.
		if exists {
			fmt.Println("This id is already used!")
			menu()
		}

		fmt.Print("Please insert Employee's Name: ")
		fmt.Scanln(&name)

		fmt.Print("Please insert Employee's Last Name1: ")
		fmt.Scanln(&lastName1)

		fmt.Print("Please insert Employee's Last Name2: ")
		fmt.Scanln(&lastName2)

		fmt.Print("Please insert Employee's Birth Date (dd/MM/yyyy): ")
		fmt.Scanln(&strBirthDate)
		birthDate, err = time.Parse("02/01/2006", strBirthDate)
		// If an date enter cannot be convert into a date with dd/MM/yyyy format returns to the menu.
		if err != nil {
			fmt.Println("Something was wrong, plis check your enter.")
			menu()
		}

		new := insertEmployee(id, name, lastName1, lastName2, birthDate)
		fmt.Println(new)

		menu()
	case 2:
		var strId string
		fmt.Print("\nPlease insert Employee's id: ")
		fmt.Scanln(&strId)
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		fmt.Println(getEmployeeAgeById(id))

		menu()
	case 3:
		fmt.Println(getAllOrderByLastName())

		menu()
	case 4:
		fmt.Println(getAllOrderByAge())

		menu()
	default:
	}
}

func insertEmployee(id int, name, lastName1, lastName2 string, birthDate time.Time) Employee {
	// To make an employee instance.
	new := newEmployee(id, name, lastName1, lastName2, birthDate)

	// Save the new employee on Employees Map.
	Employees[id] = new

	return new
}

func getEmployeeAgeById(id int) int {
	employee := getEmployeeById(id)

	// I use a library to calculate Age given a date.
	employeeAge := age.CalculateToNow(employee.BirthDate)

	return employeeAge
}

// F R O M  H E R E
type lessFunc func(p1, p2 *Employee) bool

// multiSorter implements the Sort interface, sorting the employees within.
type multiSorter struct {
	employees []Employee
	less      []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(employees []Employee) {
	ms.employees = employees
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.employees)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.employees[i], ms.employees[j] = ms.employees[j], ms.employees[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other). Note that it can call the
// less functions twice per call. We could employee the functions to return
// -1, 0, 1 and reduce the number of calls for greater efficiency: an
// exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.employees[i], &ms.employees[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

// T O  H E R E  W A S  T A K E N  F R O M  A  G O 'S  D O C U M E N T A C I O N  E X A M P L E.
// The main idea it to use one or more funtions to sort a slice in our case.

// Funtion to delete acute from a string
func removeAcutes(input string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	result, _, _ := transform.String(t, input)

	return result
}

func getAllOrderByLastName() []Employee {
	// Closures that order the Employee structure.
	// These closures have to match with the alias lessFunc.
	lastName1 := func(p1, p2 *Employee) bool {
		return removeAcutes(p1.LastName1) < removeAcutes(p2.LastName1)
	}
	lastName2 := func(p1, p2 *Employee) bool {
		return removeAcutes(p1.LastName2) < removeAcutes(p2.LastName2)
	}

	// Making a employees slice and full it with Employees map information.
	var employees []Employee
	for _, employee := range Employees {
		employees = append(employees, employee)
	}

	OrderedBy(lastName1, lastName2).Sort(employees)

	return employees
}

func getAllOrderByAge() []Employee {
	// Closures that order the Employee structure by Age. As you can see I use getEmployeeAgeById to get Age given a date.
	age := func(p1, p2 *Employee) bool {
		return getEmployeeAgeById(p1.ID) < getEmployeeAgeById(p2.ID)
	}

	var employees []Employee
	for _, employee := range Employees {
		employees = append(employees, employee)
	}

	OrderedBy(age).Sort(employees)

	return employees
}
