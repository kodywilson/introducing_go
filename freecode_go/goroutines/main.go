package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	// wg.Add(1)
	// m.RLock()
	// go sayHello()

	// var msg = "World"
	// wg.Add(1)
	// m.RLock()
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// msg = "Goodbye"
	// more advanced example
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go incremement()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func incremement() {
	counter++
	m.Unlock()
	wg.Done()
}
