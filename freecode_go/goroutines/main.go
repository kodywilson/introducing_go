package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	wg.Add(1)
	go sayHello()
	time.Sleep(100 * time.Microsecond)

	var msg = "World"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	// more advanced example
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go incremement()
	}
	wg.Wait()
}

func sayHello() {
	m.RLock()
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func incremement() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
