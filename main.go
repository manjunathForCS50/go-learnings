package main

import (
	"fmt"
	// "io"
	"os"

	// "github.com/spf13/cobra"
)

func main() {
	fmt.Println("Hello, world!");
	scanName()
}

func scanName() {
	fmt.Print("Enter your name: ")
	var name string
	_, error := fmt.Scan(&name)
	if (error != nil) {
		fmt.Println("There is an error:", error)
	}
	fmt.Println("Hello, ", name)
	os.WriteFile("./log.txt", []byte(name), 0644)
	// strWrtiter := io.StringWriter(os.Stdout)
	// cobra.WriteStringAndCheck(strWrtiter, name)
}