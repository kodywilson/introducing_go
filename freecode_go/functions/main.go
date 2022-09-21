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
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)
	// return more than one value
	d, err := divide(5.0, 0.1) // set second value to 0.0 to see error handling
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	// anonymous functions (run once)
	func() {
		fmt.Println("Hey Jude")
	}()
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
func sum(values ...int) int { // can also return with (result int)
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result // if you return with (result int) you can drop result
	// in Go, you can return pointers to local variables
	// return &result - return type would be *int
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
