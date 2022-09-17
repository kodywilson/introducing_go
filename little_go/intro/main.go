package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])

	power := getPower()
	fmt.Println(power)
}

func getPower() int {
	return 9001
}