package main

import (
	"fmt"
)

func main() {
	// Defining an empty interface and setting it to 0 integer
	// any type in go implements this interface, so it might by useful
	var i interface{} = 0
	// A switch over the type
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is an string")
	default:
		fmt.Println("I don't know what i is")
	}
}
