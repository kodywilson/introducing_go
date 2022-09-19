package main

import (
	"fmt"
)

const a int16 = 27

func main() {
	const a int = 42
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}
