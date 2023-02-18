package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// The channel is typed, hence can only process int-s
	ch := make(chan int)
	wg.Add(2)

	// This routine is receivning data from channel, implemented as an anonymous func
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	// This routine is sending data to a channel
	go func() {
		i := 42
		// We are passing a copy of data
		ch <- i
		// So this does nothing for the channel
		i = 27
		wg.Done()
	}()
	wg.Wait()
}
