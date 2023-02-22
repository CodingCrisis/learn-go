package main

import (
	"fmt"
	"time"
)

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

var logCh = make(chan logEntry, 50)

// A signal only channel - no data passed
// using empty string to achieve zero memory allocated
var doneCh = make(chan struct{})

func main() {
	// Start a simple logging function
	// This one does not enable proper cleanup
	//go logger()

	// At the end of the main function the channel used for logging is forcibly closed
	// To remedy this, we could do a defer like that:
	//defer func() {
	//	close(logCh)
	//}()

	// Another solution is to use one more channel to ask for closing the logger

	go betterLogger()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	time.Sleep(500 * time.Millisecond)
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

	// passing an empty struct, signalling end of logging
	doneCh <- struct{}{}

	// It seems to me, that te closing does not always happen in time
	// hence no "Closing the channel" info from the betterLooger proc
	// after adding a wait, it seems to always work.
	fmt.Scanln()
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

func betterLogger() {
	for {
		// Statement is blocked until a message is available in one of the channels
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("Closing the channel")
			return
			// If default is used, the select will not be blocking anymore
			//dafault: dosomething()
		}
	}
}
