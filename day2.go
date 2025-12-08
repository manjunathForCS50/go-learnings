package main 

import (
	"fmt"
	"math/rand/v2"
)

func dayTwoMain() {
	fmt.Println("----------------------------Welcome to the Number Guessing Game!----------------------------")
	fmt.Println("Enter difficult level. Choose between 1, 2, 3")
	var diffLevel int
	fmt.Scan(&diffLevel)

	for diffLevel < 1 || diffLevel > 3 {
		fmt.Println("Select between 1, 2 OR 3")
		fmt.Scan(&diffLevel)
	}

	var maxTries int

	switch diffLevel {
		case 1: 
			maxTries = 10
		case 2: 
			maxTries = 7
		case 3: 
			maxTries = 5
	}
	

	var userGuess int
	fmt.Println("Enter your guess number between 1 to 100: ")
	
	fmt.Scan(&userGuess)

	fmt.Printf("You have selected %d as your guess\n", userGuess)

	var randomNumber = rand.IntN(100) + 1

	guessCount := 0

	for userGuess != randomNumber {
		guessCount++
		if guessCount > maxTries {
			fmt.Printf("Sorry! You've used up all your tries! The correct number was %d", randomNumber)
			break
		}
		if (userGuess > randomNumber) {
			fmt.Println("Guess is too high")
			
		} else {
			fmt.Println("Guess is too low")
		}
		fmt.Println("Enter a guess again")
		fmt.Scan(&userGuess)
	}
	fmt.Printf("Correct! You guessed it in %d tries!\n", guessCount)
	fmt.Println("-----------------------------Thank you for playing!----------------------------")
}