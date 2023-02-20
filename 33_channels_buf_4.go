package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// This is a buffered channel, 50 ints long
	// Those are useful e.g. when the frequency of the sender and receiver is different
	ch := make(chan int, 50)
	wg.Add(2)

	// Func to read from channel
	go func(ch <-chan int) {
		// Iterates over a channel, as long as it's open
		for i := range ch {
			fmt.Println(i)
		}

		// Manual alternative to the for range would be
		// the comma-ok syntax enables checking if the channel is open or closed
		//for {
		//		if i, ok := <-ch; ok {
		//			fmt.Println(i)
		//		} else {
		//			break
		//		}
		//	}
		wg.Done()
	}(ch)

	// Func to write to channel
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		// Without closing the channel, the for range loop in the receiver would be deadlocked
		// as it would try to acquire new messages, when there is none sent
		// Warning - it's not possible to check if a channel is closed
		// when sending messages (only panic/recover handling)
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
