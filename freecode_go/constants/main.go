package main

import (
	"fmt"
)

// iota is scoped to constant block
const (
	x = iota
	y = iota
	z = iota
)

const (
	_  = iota // ignore first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	const a int = 42
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v, %v, %v\n", x, y, z)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
