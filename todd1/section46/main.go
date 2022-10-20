package main

import "fmt"

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter: ", <-c3+<-c4)

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	ch := factorial(4)
	for n := range ch {
		fmt.Println(n)
	}

}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func factorial(n int) chan int {
	chanOut := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		chanOut <- total
		close(chanOut)
	}()
	return chanOut
}

// func pullFact(chan int) chan int {

// }