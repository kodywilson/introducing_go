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

	// complex numbers
	var m complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", real(m), real(m))
	fmt.Printf("%v, %T\n", imag(m), imag(m))
	var h complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", h, h)

	// strings
	s := "this is a string"
	s2 := "this is a also a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Println(len(s))
	fmt.Printf("%v, %T\n", s[0], s[0])
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// runes
	r := 'a' // same as var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
