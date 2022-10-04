package main

import "fmt"

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
}
