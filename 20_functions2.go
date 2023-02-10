package main

import (
	"fmt"
)

func main() {
	// Anonymous function with direct execution
	func() {
		fmt.Println("Hello Marcin!")
	}()

	// This works but might fail in case of async execution
	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	// So this is better
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// Assigning anonymous function to a variable
	f := func() {
		// The full type definition would be:
		// var f func() =
		fmt.Println("Hello again Marcin!")
	}
	// Running the function
	f()

	// Full definition
	// Note that this function is available in this scope  and only below its declaration
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
