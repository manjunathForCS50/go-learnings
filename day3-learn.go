package main

import (
	"fmt"
)

func daythreeLearn() {
	var name string
	var age int
	var isActive bool

	fmt.Println("The name is: ", name)
	fmt.Println("The age is: ", age)
	fmt.Println("The active status is: ", isActive)

	fmt.Println("Name Address :", &name, "Age:", &age, "Active:", &isActive)
	fmt.Println("Name Address :", *&name, "Age:", *&age, "Active:", *&isActive)

	var agePtr *int = &age
	var namePtr = &name
	fmt.Println("age pointer address is: ", agePtr)
	fmt.Println("the age address is: ", &age)
	fmt.Println("the nsmre address is: ", namePtr)

	var score = 15
	var scorePtr = &score

	fmt.Println("the value of score is : ", score)

	*scorePtr = *scorePtr + 10

	fmt.Println("the new score is: ", score)
	fmt.Println("the value at score ptr address is: ", *scorePtr)
}