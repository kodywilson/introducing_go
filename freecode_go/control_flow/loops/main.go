package main

import (
	"fmt"
)

func main() {
	// simple loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// can have more than one variable in initializer
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
	// counter initialized outside loop (still need the first ;)
	counter := 0
	for ; counter < 5; counter++ {
		fmt.Printf("counter value in the loop is %v\n", counter)
	}
	fmt.Printf("counter value outside the loop is %v\n", counter)

	// other ways to break out of loop
	for {
		if counter > 8 {
			break
		}
		fmt.Printf("Now counter is equal to %v\n", counter)
		counter++
	}

	// continue statement // only print odd numbers in this example
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
