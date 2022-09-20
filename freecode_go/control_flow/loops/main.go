package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i, j := 0, 0; i < 5 && j <= 4; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even!\n", i)
		} else {
			fmt.Printf("%v is odd!\n", i)
		}
	}
}
