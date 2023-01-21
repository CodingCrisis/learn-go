package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 1
	var y float64 = 2.5
	var z uint16 = math.MaxInt16

	// Can't use operators on mismatched types, need to cast
	x = x + int(z)

	// %T gives type of variable
	fmt.Printf("%T  %v, %T %v, %T, %v\n", x, x, y, y, z, z)

	var f float32 = float32(x)
	fmt.Printf("%T %f\n", f, f)

	// Can initialize bool with a
	var b bool = 1 == 2
	fmt.Printf("%T %v\n", b, b)

	// Try some bit operators (not available on floating point)
	b1 := 10 // 1011
	b2 := 3  // 0011

	fmt.Printf("%b\n", b1&b2)  // 0010
	fmt.Printf("%b\n", b1|b2)  // 1011
	fmt.Printf("%b\n", b1^b2)  // 1001
	fmt.Printf("%b\n", b1&^b2) // 0100

	// Shift like in
	shift := 8                              // 0001000
	fmt.Println("shift 8 << 3: ", shift<<3) // 1000000
	fmt.Println("shift 8 >> 3: ", shift>>3) // 0000001

	n := 3.14 // Default is float64
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T", n, n)

	// Complex types - oh come on ;)
	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	var cfloat complex128 = complex(5.2, 12.5)
	fmt.Printf("%v, %T\n", cfloat, cfloat)
}
