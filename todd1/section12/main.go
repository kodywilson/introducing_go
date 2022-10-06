package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	bob := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	chuck := person{
		first: "Chunk",
		last:  "Norris",
		age:   72,
	}
	fmt.Println(bob, chuck)
	fmt.Println(bob.first, bob.last)
}
