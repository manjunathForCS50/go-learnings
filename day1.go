package main

import (
	"fmt"
)

func dayOneMain() {
	// var wih explicit type
	var name string
	var age int
	var isLearning bool

	fmt.Println("Printing zero values:")
	fmt.Println("name: ", name)
	fmt.Println("age:", age)
	fmt.Println("learning:", isLearning)

	// var with values
	var city string = "Bangalore"
	var exp int = 11

	company := "Docusign"
	salary := 15

	fmt.Println("\nInitialised values:")
	fmt.Println("city", city)
	fmt.Println("exp", exp)
	fmt.Println("company", company)
	fmt.Println("salary", salary)

	//constants
	const pi = 3.14
	const maxRetries = 5

	fmt.Println("\nConstanbts")
	fmt.Println("pi", pi)
	fmt.Println("maxRetries", maxRetries)
}