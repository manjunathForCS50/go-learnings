package main

import (
	"fmt"
	// "io"
	"os"
	// "github.com/spf13/cobra"
)

func main() {
	fmt.Println("Hello, world!")
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
	util()
	utilOther()
	handleArray()
	dayOneMain()
	dayTwoMain()
	daythreeLearn()
	var a = 10
	var b = 15
	swapValues(&a, &b)
	addTen(&a)
	var newBankAcc = BankAccount{accNo: "12345", balance: 3500, owner: "John Doe"}
	var initBal = newBankAcc.GetBalance()
	fmt.Println("The initial balance is: ", initBal)
	newBankAcc.Deposit(200)
	var withdrawresult1 = newBankAcc.Withdraw(50)
	var withdrawresult2 = newBankAcc.Withdraw(50000)
	fmt.Println("Withdraw result 1:", withdrawresult1)
	fmt.Println("Withdraw result 2:", withdrawresult2)
	var finalBal = newBankAcc.GetBalance()
	fmt.Println("The final balance is: ", finalBal)
	// strWrtiter := io.StringWriter(os.Stdout)
	// cobra.WriteStringAndCheck(strWrtiter, name)
}