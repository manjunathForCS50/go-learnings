package main

import (
	"fmt"
)

func util() {
	fmt.Println("In the util function")
}

func utilOther() {
	fmt.Println("Some other function")
}

func handleArray(){
	strArray:= []string{"abc", "def"}
	for index, value := range strArray {
		fmt.Println("Index:", index, "Value:", value)
	}

	for _, value := range strArray {
		fmt.Println("Printing only the value: ", value)
	}
}
