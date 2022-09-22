package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	fmt.Println("Happy Birthday", p.firstName, "!")
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(newName string) {
	p.lastName = newName
}

func main() {
	// Create a person using struct
	person1 := Person{
		firstName: "Bob",
		lastName:  "Walker",
		city:      "Atlanta",
		age:       22,
	}

	// print whole struct
	fmt.Println(person1)
	// print out one attribute
	fmt.Println(person1.city)
	// change something
	person1.age = 23
	fmt.Printf("%s is now %d years old!\n", person1.firstName, person1.age)
	// Call greet method
	fmt.Println(person1.greet())
	// call pointer method for birthday
	person1.hasBirthday()
	fmt.Printf("%s is now %d years old!\n", person1.firstName, person1.age)

	// Person gets married!
	person1.getMarried("Runner")
	fmt.Println(person1)
}
