package main

import (
	"fmt"
)

func main() {
	// section 9 exercises
	//ex 1
	for i := 1; i <= 10000; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// ex2
	for i := 65; i < 91; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
	// ex3
	counter := 0
	fmt.Println("Years you have been alive")
	for i := 1977; i <= 2022; i++ {
		fmt.Println(i)
		counter++
	}
	fmt.Println("You have been alive", counter, "years!")
}
