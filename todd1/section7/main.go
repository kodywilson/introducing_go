package main

import "fmt"

const h int = 42 // typed constant
const i = 42     // untyped constant

const (
	k = 2022 + iota
	l = 2022 + iota
	m = 2022 + iota
	n = 2022 + iota
)

func main() {
	// ex1
	a := 42
	fmt.Printf("a in decimal: %d\tbinary: %b\thex: %#X\n", a, a, a)

	// ex2
	b := (42 == 42)
	c := (42 <= 43)
	d := (42 >= 43)
	e := (42 != 43)
	f := (42 < 43)
	g := (42 > 43)
	fmt.Println(b, c, d, e, f, g)

	// ex3
	fmt.Printf("%T %T\n", h, i)

	// ex4
	h := 42
	fmt.Printf("%d\t%b\t%#x\n", h, h, h)
	i := h << 1
	fmt.Printf("%d\t%b\t%#x\n", i, i, i)

	// ex5
	j := `here is a raw string literal
	which can contain "quotes" and other things
	including line breaks`
	fmt.Println(j)

	// ex6
	fmt.Println(k, l, m, n)
}
