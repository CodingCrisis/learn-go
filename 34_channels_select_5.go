package main

import (
	"fmt"
	"sync"
	"time"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

func main() {
	go logger()

}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2023-02-22T23:47:00"), entry.severity, entry.message)
	}
}
