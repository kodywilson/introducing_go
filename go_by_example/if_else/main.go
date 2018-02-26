package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is evenly divisible by 4")
	}
	for num := -10; num < 11; num++ {
		if num < 0 {
			fmt.Println(num, "is negative")
		} else if num < 10 {
			fmt.Println(num, "has 1 digit")
		} else {
			fmt.Println(num, "has multiple digits")
		}
	}
}
