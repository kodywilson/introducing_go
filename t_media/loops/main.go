package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d\t", i)
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Print("\n")
		}
		// if else works too
		// if i % 3 == 0 && i % 5 == 0 {
		// 	fmt.Println("FizzBuzz")
		// } else if i % 3 == 0 {
		// 	fmt.Println("Fizz")
		// } else if i % 5 == 0 {
		// 	fmt.Println("Buzz")
		// }
	}
}