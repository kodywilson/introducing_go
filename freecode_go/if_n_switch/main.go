package main

import "fmt"

func main() {
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
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
