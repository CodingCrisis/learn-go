package main

import (
	"fmt"
)

func main() {
	// There is no automatic fall-through in golang (see 3rd example, it uses the fallthrouhg keyword)
	// Breaks are implied
	// But multiple values in a condition are possible
	switch 4 {
	case 1, 5, 10:
		fmt.Println("one")
	case 2, 4, 6:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	// Defining variables (with assignment) is possible
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one")
	case 2, 4, 6:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	// There is also a tagless syntax
	// It enables overlapping of variables
	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case j <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twnety")
	}

	// You can do a type switch
	var k interface{} = "2.5"
	switch k.(type) {
	case int:
		fmt.Println("k is an int")
	case float64:
		fmt.Println("k is a float64")
	case string:
		fmt.Println("k is a string")
		// We can use break if it is necessary for some specific reason
		// E.g. validating data before doing something and breaking if the validation fails
		break
		fmt.Println("and do I print this or not?")
	case [2]int:
		fmt.Println("An integer array of size 2")
	default:
		fmt.Println("k is another type")
	}
}
