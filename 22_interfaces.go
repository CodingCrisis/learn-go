package main

import (
	"fmt"
)

func main() {
	// Anything that implements the interface can be assigned to the object
	// The interface implementation in GO is implicit
	// So one could define interfaces for existing methods
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Marcin!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// Interface does not describe data, but behavior
// The convention for a single method interface is adding "er" to the name of the function
// Hence Write and Writer
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// this is a method of the type above
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

// As you can't add methods directly to the int type, let's do an alias
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
