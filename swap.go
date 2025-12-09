package main

import (
	"fmt"
)

func swapValues(a, b *int) {
	fmt.Println("We''\re swapping values")

	var tempAddress = *a
	*a = *b
	*b = tempAddress

	fmt.Println("After swapping inside function a:", *a, " b:", *b)
}