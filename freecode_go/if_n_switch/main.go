package main

import (
	"fmt"
	"math"
)

func main() {
	// if statements
	if true {
		fmt.Println("The test is true")
	}
	number := 50
	guess := 30
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
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
		fmt.Println(number <= guess, number >= guess, number != guess)
		fmt.Println(returnTrue())
	}
	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.0001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// switch statements
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
