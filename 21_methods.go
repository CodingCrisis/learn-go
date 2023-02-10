package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Marcin",
	}
	g.greet()
	fmt.Println("The new name is: ", g.name)

	g.greet2()
	fmt.Println("The new name is: ", g.name)
}

type greeter struct {
	greeting string
	name     string
}

// This is a method
// Just a function, that is executed in a known context
// So in this case the (g greeter) part
// It's a value receiver, so the function works on a copy of the object
// Hence the name change only works in scope of thie greet function
// It also means that it could be costly on big objects
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Stefan"
}

// It's possible to pass via reference, using a pointer receiver
// Go uses here the implicit dereferencing of pointers, so no need to use &
func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = "Stefan"
}
