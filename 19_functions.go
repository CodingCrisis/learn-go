package main

import (
	"fmt"
)

func main() {
	sayMessage("Hello from function!")
	for i := 0; i < 5; i++ {
		sayAnotherMessage("Hello from another function. ", i)
	}
	sayGreeting("Hi", "Marcin")

	greeting := "Hello"
	name := "Marcin"
	sayGreetingValue(greeting, name)
	fmt.Println("Has name changed when passing to function via value?", name)

	sayGreetingRef(&greeting, &name)
	fmt.Println("Has name changed when passing to function via value?", name)
}

// Upper case name is public, lowercase not
func sayMessage(msg string) {
	fmt.Println(msg)
}

// multiple params
func sayAnotherMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is: ", idx)
}

// if parameters are the same type, just list them and put type at the end
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

// You can pass by value
func sayGreetingValue(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Stefan"
	fmt.Println("Name in function:", name)
}

// You can pass by reference
// which is often more performant (e.g. passing a bit structure)
// maps/slices are ref types, so there's no choice
func sayGreetingRef(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Stefan"
	fmt.Println("Name in function:", *name)
}
