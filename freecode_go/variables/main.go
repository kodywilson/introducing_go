package main

import (
	"fmt"
	"strconv"
)

// cannot use short form to declare variables at this scope level

var (
	a int     = 3
	b int     = 4
	c float32 = 5.12
)

func main() {
	var i int
	fmt.Println(42)
	i = 5
	fmt.Println(i)
	i = 25
	fmt.Println(i)
	x := 10000
	fmt.Println(x)
	var y int = 2000
	fmt.Println(y)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Println(a * b)
	fmt.Println(a + int(c))
	fmt.Println(strconv.Itoa(a))
}
