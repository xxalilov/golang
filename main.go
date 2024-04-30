package main

import (
	"fmt"
	"time"
)

func main() {
	// Print a welcome message
	fmt.Println("Hello! What is your name?")

	// Reading input from user
	var name string
	fmt.Scanln(&name)

	// Greet the user
	fmt.Printf("Nice to meet you, %s!\n", name)

	// Ask for the user's birth year
	fmt.Println("What year were you born in?")

	// Read the birth year
	var birthYear int
	fmt.Scanln(&birthYear)

	// Calculate the age
	currentYear := time.Now().Year()
	age := currentYear - birthYear

	// Display the age
	fmt.Printf("%s, you are %d years old.\n", name, age)
}
