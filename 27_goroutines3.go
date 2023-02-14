package main

import(
	"fmt"
	"sync"
)

// This is a wait group used to sync gorutines
var wg = sync.WaitGroup{}

func main(){
	var msg = "Hello"
	wg.Add(1)
	go func(msg string){
		fmt.Println( msg)
		// Notifies wait group that routine is done
		wg.Done()
	}(msg)
		
	msg = "Goodbye"

	// Waits on the go routines to finish
	wg.Wait()
}
