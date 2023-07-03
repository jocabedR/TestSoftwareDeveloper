package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("Incorrect number of arguments.\n")
		os.Exit(1)
	}

	characters := os.Args[1]

	log.Println("Total of founded numbers: ", countNumbers(characters))
}

func countNumbers(characters string) (count int) {
	// I decided to use a Regular Expression to resolve the problem.
	regex := regexp.MustCompile(`\d`)
	// -1 argument make us sure that we are looking for unlimited number of matches within the string.
	matches := regex.FindAllString(characters, -1)

	return len(matches)
}
