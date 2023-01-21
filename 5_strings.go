package main

import (
	"fmt"
)

func main() {
	// Strings are UTF8 chars
	s := "This is a string"
	s2 := "! :)"

	fmt.Printf("%v, its type is: %T\n", s, s)
	// Can access string as a table
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))

	// Can concat strings
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// Can interpret as bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// Runes are UTF32 chars, actualy it's an int32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

	r_array := []rune(s)
	fmt.Printf("Array of runes: %U\n", r_array)
}
