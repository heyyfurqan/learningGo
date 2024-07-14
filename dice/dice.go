//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	fmt.Println("Welcome to dawg-roller!")

	rand.Seed(time.Now().UnixNano())

	var rolls int
	fmt.Println("How many times do you want to roll the dice?")
	fmt.Scan(&rolls)

	var dice int
	fmt.Println("How many number of dice you want to use for rolling?")
	fmt.Scan(&dice)

	var sides int
	fmt.Println("How many number of sides of the dice?")
	fmt.Scan(&sides)

	totalRoll := 0

	for r := 1; r <= rolls; r++ {
		for d := 1; d <= dice; d++ {
			currentRoll := roll(sides)
			fmt.Println("Current rolled number is: ", currentRoll)
			totalRoll += currentRoll
		}
	}
	fmt.Println("Total roll is: ", totalRoll)

	if totalRoll == 2 && dice == 2 {
		fmt.Println("Snake eyes")
	} else if totalRoll == 7 {
		fmt.Println("Lucky 7")
	} else if totalRoll%2 == 2 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
