package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// os.Args has to content the file the user it's running and the date they want to analyze.
	if len(os.Args) != 2 {
		log.Print("Incorrect number of arguments.\n")
		os.Exit(1)
	}

	stringDate := os.Args[1]

	validateDate(stringDate)
}

func validateDate(stringDate string) (isValid bool, date time.Time) {
	var err error
	// using January 2nd, 2006 as a reference date in the time package in Go
	// is simply a convention established by the designers of the Go language.
	// This specific date was chosen because it allows for an easy and unambiguous representation of date formats.
	date, err = time.Parse("02/01/2006", stringDate)

	//
	if err != nil {
		isValid = false
		log.Println("Invalid date!:", err)
		return
	}

	isValid = true
	log.Println("Valid date!")
	return
}
