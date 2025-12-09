package main

import (
	"fmt"
)

func addTen(a *int) {
	*a = *a + 10
	fmt.Println("the value at a was chanmged from inside function to: ", *a)
}