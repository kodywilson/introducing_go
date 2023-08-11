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
	// pointers with structs
	var as aStruct
	as = aStruct{foo: 300}
	fmt.Println(as)
	var ms *myStruct
	ms = new(myStruct) // cannot use object initialization syntax with new
	// Don't have to use (*ms).foo - Go allows ms.foo
	ms.foo = 42 // the compiler will derefence for you
	fmt.Println(ms.foo)

	// pointers with arrays and slices
	// arrays get copied while slices point to original slice
	x := []int{1, 2, 3}
	y := x
	fmt.Println(x, y)
	x[1] = 42
	fmt.Println(x, y)

	// maps also point to the original map when copied, like slices
	c := map[string]string{"foo": "bar", "baz": "buz"}
	d := c
	fmt.Println(c, d)
	c["foo"] = "qux"
	fmt.Println(c, d)

	// quick recap, primitives, arrays, and structs copy the object into new memory
	// maps and slices, copy the memory location so altering the original, alters the copy as well
}

type aStruct struct {
	foo int
}

type myStruct struct {
	foo int
}
