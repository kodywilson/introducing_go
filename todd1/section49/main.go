package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
	fmt.Println("This is a test. I repeat, this is a test.")
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	}
	return fmt.Sprint("x is less than 10")
}
