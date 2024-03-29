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

	// variatic parameter
	sum("And the sum is: ", 1, 2, 3, 4, 5)

	s2 := sum2(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", s2)

	s3 := sum3(1, 2, 3, 4, 5)
	fmt.Println("Sum again, but reference:", *s3)

	s4 := sum4(1, 2, 3, 4, 5)
	fmt.Println("Sum again, but named return param:", s4)

	// Typical usage of err return parameter
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

// Upper case name is public, lowercase not
func sayMessage(msg string) {
	fmt.Println(msg)
}

// multiple params
func sayAnotherMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is:", idx)
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

// Function with multiple input variables
// variatic parameter can be only one per function and it needs to be the last one
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	// variatic params are available as a slice
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

// Same as "sum", but with an output value (of int)
func sum2(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// Same as "sum2" but result returned via reference
// This is specific to Golang, the value is going to be defined on the heap not on the stack
// so it will not be removed on finishing the function
func sum3(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// Same as "sum", but with named return parameter
// Named parameter is zero initialized
// apparently this is not very popular in go, as could be confusing
func sum4(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// Example of returning multiple values from a function
func divide(a, b float64) (float64, error) {
	// we could panic here:
	// if b == 0.0 {
	// panic("Cannot provide zero as second value")
	// }

	// but we could return an error via a second parameter
	if b == 0.0 {
		// Errorf provides a possibility to create error messages
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
