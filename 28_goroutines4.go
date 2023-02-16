package main

import (
	"fmt"
	"runtime"
	"sync"
)

// This is a wait group used to sync gorutines
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// Make sure it's possible to create many threads
	runtime.GOMAXPROCS(100)

	// Create 20 go routines in a loop
	// The routines are racing each other without any wait group of mutex
	// That's actually a dumb example, as the gorutines are waiting on locked resource
	// there is no actual concurrency
	for i := 0; i < 10; i++ {
		wg.Add(2)
		// A read mutex - infinite number of readers, but only one writer
		// when writer is doing it's thing, readers cannot access the locked resource
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
