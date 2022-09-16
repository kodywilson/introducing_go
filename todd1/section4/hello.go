package main

import "fmt"

func main() {
	x := 42
	fmt.Println("Hello Everyone!")
	foo()
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	fmt.Println("something more")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
