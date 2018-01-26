package main

import "fmt"

func main() {
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Println(i, "Evenly divided by 3!")
		}
	}
	// FizzBuzz!
	for i := 1; i < 201; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, "FizzBuzz")
				continue
			} else {
				fmt.Println(i, "Fizz")
			}
		}
		if i%5 == 0 {
			if i%3 == 0 {
				fmt.Println(i, "FizzBuzz")
				continue
			} else {
				fmt.Println(i, "Buzz")
			}
		}
	}
}
