package main

import "fmt"

func main() {
	// Control Flow
	// loops
	// for init; condition; post {}
	for i := 0; i < 5; i++ {
		fmt.Println("i is now", i)
	}
	// nesting loops
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			//fmt.Printf("%v ", j)
			fmt.Printf("%v ", i*j)
		}
		fmt.Println()
	}

	a := 2
	b := 64
	// similar to while
	for a < b {
		a *= 2
		fmt.Println("another round")
		if a > 31 {
			break
		}
	}

	// continue
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%v is even\n", i)
	}

	// print out ascii characters
	fmt.Println("Print out decimal 33 - 127 as ascii chars")
	for i := 33; i < 127; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}
