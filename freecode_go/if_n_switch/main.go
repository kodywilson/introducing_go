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
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five, or ten")
	case 2, 4, 6:
		fmt.Println("two")
	default:
		fmt.Println("not matched")
	}
	// another way to do case statement
	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		//fallthrough - fall through to next case
		// Go has implicit breaks in case statements
	case j <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than twenty")
	}
	// switch on interface
	var k interface{} = 1
	switch k.(type) {
	case int:
		fmt.Println("k is an int")
		break // break out early
		fmt.Println("This will print too")
	case float64:
		fmt.Println("k is a float64")
	case string:
		fmt.Println("k is a string")
	default:
		fmt.Println("k is another type")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
