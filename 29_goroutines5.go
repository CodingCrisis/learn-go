package main

import (
	"fmt"
	"runtime"
)

func main() {
	// It's possible to set the maximum of available threads
	// Recommended minimum is a thread for each processor core (or more)
	// -1 means default/current setting
	runtime.GOMAXPROCS(-1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
