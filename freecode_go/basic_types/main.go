package main

import (
	"fmt"
)

func main() {
	// bools and ints
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

	// and or xor
	x := 10             // 1010
	y := 3              // 0011
	fmt.Println(x & y)  // 0010
	fmt.Println(x | y)  // 1011
	fmt.Println(x ^ y)  // 1001
	fmt.Println(x &^ y) // 0100

	// bit shifting
	z := 8              // 00001000
	fmt.Println(z << 3) // 01000000 = 64 // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(z >> 3) // 00000001 = 1  // 2^3 / 2^3 = 2^0 = 1

	// floats
	g := 3.14
	g = 13.7e72
	g = 2.1e14
	fmt.Printf("%v, %T\n", g, g)
}
