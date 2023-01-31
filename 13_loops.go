package main

import (
	"fmt"
)

func main() {
	// Regular for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Test1: ", i)
	}

	// Increment can be specific
	for i := 0; i < 5; i = i + 2 {
		fmt.Println("Test2: ", i)
	}

	// Multiple indexes is possible, but looks ugly
	// can't use i, j = i++, j++
	// because i++ is a statement, not an expression
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println("Test3: ", i, j)
	}

	// Weird things are allowed to be done to the index
	// But it does not seem to be a great practice ;)
	for i := 0; i < 5; i++ {
		fmt.Println("Test4: ", i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}

	// Index can be defined outside
	// if we need it outside of the for loop
	k := 0
	for ; k < 5; k++ {
		fmt.Println("Test5: Inside the loop: ", k)
	}
	fmt.Println("Test5: After the loop: ", k)

	/* That would be an infinite loop, semicolons are optional and get deleted by fmt
	k = 0
	for k < 5 {
		fmt.Println("Test6: ", k)
	}
	*/

	// We can manipulate the index inside the loop
	k = 0
	for k < 5 {
		fmt.Println("Test7: ", k)
		k++
	}

	// You can break loops
	k = 0
	for {
		fmt.Println("Test8: ", k)
		k++
		if k == 5 {
			break
		}
	}

	// Alternatively you can use continue to skip to next iteration of the loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Test9: ", i)
	}

	// Nested loops are fine as well
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("Test10: ", i*j)
		}
	}

	// Break works only on the nearest loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("Test11: ", i*j)
			if i*j >= 3 {
				break
			}
		}
	}

	// but there are labels - goto was bad? ;)
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("Test12: ", i*j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// Working with slices/arrays using for range
	// Always the key/value format
	s := []int{1, 2, 3}
	fmt.Println(s)
	for key, val := range s {
		fmt.Println("Test13: key: ", key, " val: ", val)
	}

	// Working with maps
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for key, val := range statePopulations {
		fmt.Println("Test14: key: ", key, " val: ", val)
	}
	// To only use the key, simply skip the second variable

	// To only use the value, use the underscore operator like this
	/*
		for _, val := range statePopulations {
			fmt.Println("Test14: key: ", key, " val: ", val)
		}
	*/

	// Or even iterating a string
	str := "hello Go!"
	for key, val := range str {
		fmt.Println("Test15: key: ", key, " val: ", string(val))
	}
}