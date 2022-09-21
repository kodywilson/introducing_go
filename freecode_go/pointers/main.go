package main

import "fmt"

func main() {
	var a int = 42
	// use * to create pointer, set type to same as target variable
	var b *int = &a    // use & (address of operator) to get memory address
	fmt.Println(&a, b) // show memory addresses
	// now use dereferencing operator to get value at pointer
	fmt.Println(a, *b) // show memory addresses
	a = 27
	fmt.Println(a, *b)
	*b = 77
	fmt.Println(a, *b)
	fmt.Printf("The type of b is %T\n", b)
	// Go does not allow pointer arithmetic without using unsafe
	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42 // the compiler will derefence for you
	fmt.Println(ms.foo)

	c := map[string]string{"foo": "bar", "baz": "buz"}
	fmt.Println(c)
}

type myStruct struct {
	foo int
}
