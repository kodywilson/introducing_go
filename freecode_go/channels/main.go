package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

func main() {
	go logger()
	// defer func() { // can use this to close channel
	// 	close(logCh)
	// }()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	// channels are always made with make()
	ch := make(chan int, 50)
	//for j := 0; j < 5; j++ {
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		//fmt.Println(len(ch))
		wg.Done()
	}(ch)
	//}
	wg.Wait()

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} // shut the channel down
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}

	}
}
