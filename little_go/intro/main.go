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

	name, str := "Bob", 18
	fmt.Printf("%s's strength is %d\n", name, str)
	numby, booly := twoReturn(20)
	fmt.Println(numby, booly)
}

func getPower() int {
	return 9001
}

func twoReturn(val int) (int, bool) {
	return val * 2, true
}