package main

import (
	"fmt"
)

type BankAccount struct {
	accNo string
	balance float32
	owner string
}

func (ba *BankAccount) Deposit(amount float32) {
	ba.balance = ba.balance + amount
	fmt.Println("Dear ", ba.owner, " the amount ", amount, " has been sucessfully added to the account: ", ba.accNo)

}

func (ba *BankAccount) Withdraw(withdrawAmt float32) bool {
	if withdrawAmt < 0 || withdrawAmt > ba.balance {
		fmt.Println("The entered withdrawal ", withdrawAmt ," amount is invalid")
		return false
	}

	ba.balance = ba.balance - withdrawAmt
	fmt.Println("Dear ", ba.owner, " the amount ", withdrawAmt, " has been sucessfully withdrawn from the account: ", ba.accNo)
	return true
}

func (ba BankAccount) GetBalance() float32 {
	return ba.balance
}