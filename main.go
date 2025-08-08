package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

type Level struct {
	name  string
	tries int
}

func main() {
	var levels = [3]Level{{name: "Easy", tries: 10}, {name: "Medium", tries: 5}, {name: "Hard", tries: 3}}
	var selectedLevel int8
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Please select the difficulty level:")

	for index, level := range levels {
		fmt.Printf("%v. %v (%v chances)\n", index+1, level.name, level.tries)
	}

	fmt.Printf("Enter your choice: ")
	_, err := fmt.Scanln(&selectedLevel)
	if err != nil || selectedLevel < 1 || selectedLevel > 3 {
		log.Fatal("Invalid level selection...")
	}

	fmt.Printf("Great! You have selected the %v difficulty level\n", levels[selectedLevel-1].name)

	// generate the random number
	toGuess := rand.IntN(100)
	guessed := -1
	tries := 0
	for toGuess != guessed && tries < levels[selectedLevel-1].tries {
		fmt.Printf("Enter your guess: ")
		_, err = fmt.Scanln(&guessed)

		tries++

		if err != nil {
			fmt.Println("Incorrect guess!")
			continue
		}

		if guessed == toGuess {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.\n", tries)
			break
		}

		var moreOrLess string
		if guessed < toGuess {
			moreOrLess = "greater"
		} else {
			moreOrLess = "less"
		}
		fmt.Printf("Incorrect! The number is %v than %v.\n", moreOrLess, guessed)
	}

	if tries == levels[selectedLevel-1].tries {
		fmt.Printf("Loose! Maybe you will be luckiest next time.")
	}
}
