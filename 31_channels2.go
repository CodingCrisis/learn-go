package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// By default the channels are unbuffered (only 1 message at a time allowed)
	ch := make(chan int)

	// Create 5 senders and 5 receivers
	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}
