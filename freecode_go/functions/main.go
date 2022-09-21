package main

import (
	"fmt"
)

func main() {
	sayMessage("Hello Go")
}

func sayMessage(msg string) {
	fmt.Println(msg)
}
