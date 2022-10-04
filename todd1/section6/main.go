package main

import "fmt"

func main() {
	// boolean type - true / false
	var booly bool
	fmt.Println("booly is", booly)

	a := 7
	b := 42
	fmt.Println(a == b-35)

	// numbers - integers and floats
	x := 42
	y := 42.34534
	fmt.Printf("x type is %T while y type is %T\n", x, y)
	// int types, int8, int16, int32, int64
	// u before int signifies unsigned (positive only)
	// those follow the same pattern uint8, uint16, uint32, etc.
	var bite byte = 42
	fmt.Printf("bite is type %T\n", bite)

	// strings
	s := "Hello, people" // use `` for string literals
	fmt.Println(s)

	bs := []byte(s) // convert string to slice of bytes
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) // prints out ascii codes of each letter in the slice

	for i := 0; i < len(s); i++ {
		fmt.Printf("ASCII: %v\tUnicode: %#U\tHexidecimal: %#X\n", s[i], s[i], s[i])
	}

	// rune is an alias for int32
	// unicode uses int32 (4 bytes)
	n := bs[0]
	fmt.Println(n)
	fmt.Printf("type: %T binary: %b hex: %#X\n", n, n, n)
}
