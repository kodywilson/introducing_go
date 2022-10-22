package main

import "fmt"

func main() {
	var name string = "Bob"
	var age int32 = 37
	const isCool = true
	fmt.Println(name, age)
	fmt.Printf("name is type %T and age is type %T\n", name, age)
	fmt.Println("True or False, am I cool?", isCool)
	// Primitives
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 (unsigned, always positive, no negatives)
}