package main

import (
	"fmt"
)

func main() {
	// c := make(chan int) // unbuffered channel, blocks

	// // first run with listening function instead of range
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		fmt.Println(<-c)
	// 	}
	// }()

	// time.Sleep(time.Second)
	// fmt.Println("Done with first round")

	// // now, use range this time
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }

	// fmt.Println("Done with second round")
	// // n-1 - didn't work, trying semaphore
	// cha := make(chan int)
	// done := make(chan bool)
	// // create wait group
	// //var wg sync.WaitGroup

	// //wg.Add(2)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		cha <- i
	// 	}
	// 	//wg.Done()
	// 	done <- true
	// }()

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		cha <- i
	// 	}
	// 	//wg.Done()
	// 	done <- true
	// }()

	// go func() {
	// 	//wg.Wait()
	// 	<-done // you can throw values away off channel
	// 	<-done
	// 	close(cha)
	// }()

	// for num := range cha {
	// 	fmt.Println(num)
	// }
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
