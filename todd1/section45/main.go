package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // unbuffered channel, blocks

	// first run with listening function instead of range
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println()

	// now, use range this time
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
