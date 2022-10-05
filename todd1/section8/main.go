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

	// if then conditionals
	if true {
		fmt.Println("True")
	}
	if !false {
		fmt.Println("Not false (true)")
	}

	// local scope variable part of conditional
	if x := 42; x == 42 {
		fmt.Println("true")
	}
	// fmt.Println(x) no access to x in this case. It is
	// local to the if statement

	x := 39
	if x == 40 {
		fmt.Println("value is 40")
	} else if x > 40 {
		fmt.Println("value is over 40")
	} else {
		fmt.Println("value is under 40")
	}
}
