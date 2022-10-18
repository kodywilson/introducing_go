package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// create wait group
var wg sync.WaitGroup

// global counter
// var counter int
var counter int64

// mutex - allows to lock and unlock variables
var mutex sync.Mutex

func main() {
	wg.Add(2)
	//go foo()
	//go bar()
	go incrementor(("Foo:"))
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func foo() {
	for i := 0; i < 500; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 500; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func incrementor(s string) {
	for i := 0; i < 500; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		//mutex.Lock()
		//x := counter
		//x++
		//counter = x
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
		//mutex.Unlock()
		// could replace most of the above by using atomic
		// atomic.AddInt64(&counter, 1)
	}
	wg.Done()
}
