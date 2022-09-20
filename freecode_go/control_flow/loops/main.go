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

	// nested loops
	fmt.Println("Now print nested loop results")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}

	// looping over collections
	s := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Println(s[i])
	}
	for k, v := range s {
		fmt.Println(k, v)
	}
	// over maps also
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	// over strings
	stringy := "Hello Go!"
	for k, v := range stringy {
		fmt.Println(k, v, string(v))
	}
}
