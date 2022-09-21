package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	sayMessage("Hello Go!", i)
	// }
	greeting := "Hello"
	name := "Bob"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
	sum(1, 2, 3, 4, 5)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// pass in pointer to work with that variable directly (not a copy)
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// variadic parameters must be the last one passed in and only one is allowed
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}
