package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// Basically a nicer version of the first channel example
	// Where the functions clearly state if they read or write to a channel

	ch := make(chan int)
	wg.Add(2)

	// Func to read from channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// Func to write to channel
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()
}
