package main

import (
	"fmt"
)

func main() {
	var t bool = true
	n := 42
	o := -42
	var p uint = 10000
	a := 1 == 1
	b := 1 == 2
	fmt.Printf("%v, %T\n", t, t)
	fmt.Println(t)
	fmt.Printf("a = %v and b = %v\n", a, b)
	fmt.Printf("n = %v and is a %T\n", n, n)
	fmt.Printf("o = %v and is a %T\n", o, o)
	fmt.Printf("p = %v and is a %T\n", p, p)
}
