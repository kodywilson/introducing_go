package main

import (
	"fmt"
	"os"
)

func main() {
	argsAll := os.Args
	argsOnly := argsAll[1:]

	fmt.Println(argsAll)
	fmt.Println(argsOnly)
}
