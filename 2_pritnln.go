package main

import (
	"fmt"
	"time"
)

func main() {
	//Just print
	fmt.Print("Hello")

	//Print + new line
	fmt.Println(" World!")

	//Print with formatting
	fmt.Printf("Today is %s\n", time.Now().Weekday())

	//Print different formats (with positional args
	fmt.Printf("Default %[1]v, decimal: %[1]d, hex: %[1]x", 17)

	//Print with named args - not sure if possible?
	//fmt.Printf("My name is [surname], [name] [surname]\n", {name = "James"}, {surname = "Bond"})
}
