package main

import (
	"fmt"
	"math"
)

func main(){
	var x int = 1
	var y float64 = 2.5
	var z uint16 = math.MaxInt16

	// Can't use operators on mismatched types, need to cast
	x = x + int(z)
			
	fmt.Printf("%T  %v, %T %v, %T, %v\n", x, x, y, y, z, z)
	
	var f float32 = float32(x)
	fmt.Printf("%T %f\n", f, f)
	
	var b bool = 1 == 2
	fmt.Printf("%T %v\n", b, b)

	b1 := 10 // 1011
	b2 := 3  // 0011

	fmt.Printf("%b\n", b1 & b2) // 0010
	fmt.Printf("%b\n", b1 | b2) // 1011
	fmt.Printf("%b\n", b1 ^ b2) // 1001
	fmt.Printf("%b\n", b1  &^ b2) // 0100
	
	n := 3.14
	n = 13.7e72
	n = 2.1E14
	fmt.Printf("v, %T", n, n) 
}
