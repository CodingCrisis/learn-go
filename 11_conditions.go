package main

import (
	"fmt"
	"math"
)

func main() {
	// Go requires curly braces with if-s
	if true {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	// The pop/ok is not available outside of the if statement
	// so the following here would fail:
	// fmt.Println(pop)

	// Guessing game
	number := 50
	guess := 30

	// the function returnTrue is only used as an example of short cricuit
	// e.g. with guess = -5 it would not be executed
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {

		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

	// There are standard logical operators OR || AND && and NOT !
	fmt.Printf("Bang logic operator: %t\n", !false)

	// Just the remaining comparison operators
	fmt.Println(number <= guess, number >= guess, number != guess)

	// nicer version with an else
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than  100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

	// Look out for IFs on floating point
	// The usual cases occur with representation and comparing
	// 0.12 is fine, but 0.123 already brakes
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// This solution is still wrong, but a bit better :)
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}

// Just a func for the condition short circuit test
func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
