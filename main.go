package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100.
	target := rand.Intn(100) + 1

	fmt.Println("Guess the number between 1 and 100!")

	var guess int
	attempts := 0

	// Loop until the guess is correct.
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You've guessed the number correctly in %d attempts.\n", attempts)
			break
		}
	}
}
