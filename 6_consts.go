package main

import (
	"fmt"
)

// External constant is a fallback to an internal one
const myConst int = 27

// Enumerated constants, iota starts at 0
const (
	_ = iota // _ throws away the value (in this case a 0)
	// usable, as not initiated int is also 0
	// so it's safer to not use 0 in enums
	a // 1
	b // 2
	c // 3
)

// Constants for data size, using bit shifting
const (
	_  = iota // ignore first
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Bit constants, a naive map of access rights
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {
	// There is no naming convention (e.g. ALL CAPS)
	// Starting from uppercase if it needs to be exported (inernal with lowercase)
	const myConst int = 42
	// must be assignable at compile time, so function doesn't cut it
	// e.g. math.Sin(1,57)
	fmt.Printf("The answer is: %v, %T\n", myConst, myConst)

	var myVariable int = 27
	fmt.Printf("Adding constants and variables is fine: %v", myConst+myVariable)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Println(KB, MB, GB, TB)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)

	// Define roles of a user as a map of bits
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}