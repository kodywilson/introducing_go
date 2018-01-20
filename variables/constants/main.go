package main

import "fmt"

const x string = "Hello Squirrel!"

func main() {
	fmt.Println(x)
	funkadelic()
}

func funkadelic() {
	const y string = "Say it again!"
	fmt.Println(y)
	fmt.Println(x)
}
