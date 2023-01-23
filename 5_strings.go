package main

import (
	"fmt"
)

func main() {
	// Strings are UTF8 chars
	s := "This is a string"
	s2 := "! :)"

	// Strings are immutable, so this would fail:
	// s[2] = 't'

	fmt.Printf("%v, its type is: %T\n", s, s)
	// Can access string as a table
	fmt.Printf("Third char of the string is: %v, of type %T\n", s[2], s[2])
	fmt.Printf("You can cast it to string: %v, of type %T\n", string(s[2]), string(s[2]))

	// Can concat strings
	fmt.Printf("Concatenated strings: %v, %T\n", s+s2, s+s2)

	// Can interpret as bytes
	b := []byte(s)
	fmt.Printf("Casted to an Array of bytes: %v, %T\n", b, b)

	// Runes are UTF32 chars, actualy it's an int32
	var r rune = 'a'
	fmt.Printf("A rune is a weird string: %v, %T\n", r, r)

	// A rune slice; it's not immutable like a string
	r_array := []rune(s)
	fmt.Printf("Array of runes: %U\n", r_array)
}
