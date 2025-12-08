package main


import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
)

func dayTwoMain() {
    fmt.Println("Welcome to the Number Guessing Game!")
    reader := bufio.NewReader(os.Stdin)

    // validate difficulty first (only accept pure digits)
    var diffLevel int
    for {
        fmt.Print("Enter difficulty level (1=easy, 2=medium, 3=hard): ")
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Input error, try again.")
            continue
        }
        token := strings.TrimSpace(line)
        n, err := strconv.Atoi(token)
        if err != nil {
            fmt.Println("Invalid input — please enter only digits.")
            continue
        }
        if n >= 1 && n <= 3 {
            diffLevel = n
            break
        }
        fmt.Println("Please select 1, 2 or 3.")
    }

    triesByLevel := map[int]int{1: 10, 2: 7, 3: 5}
    maxTries := triesByLevel[diffLevel]

    randomNumber := rand.Intn(100) + 1
    guessCount := 0

    for guessCount < maxTries {
        fmt.Printf("Attempt %d/%d - Enter your guess (1-100): ", guessCount+1, maxTries)
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Input error, try again.")
            continue
        }
        token := strings.TrimSpace(line)
        n, err := strconv.Atoi(token)
        if err != nil {
            fmt.Println("Invalid input — please enter only digits.")
            continue
        }
        if n < 1 || n > 100 {
            fmt.Println("Number out of range, enter 1-100.")
            continue
        }

        userGuess := n
        guessCount++
        if userGuess == randomNumber {
            fmt.Printf("Correct! You guessed it in %d tries!\n", guessCount)
            return
        }
        if userGuess > randomNumber {
            fmt.Println("Guess is too high.")
        } else {
            fmt.Println("Guess is too low.")
        }
        fmt.Printf("Tries left: %d\n", maxTries-guessCount)
    }

    fmt.Printf("Sorry! You've used up all your tries! The correct number was %d\n", randomNumber)
}