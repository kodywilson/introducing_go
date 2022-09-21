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
	sayGreeting(greeting, name)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}
