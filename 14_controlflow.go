package main

import (
	"fmt"
)

func def1() {
	fmt.Println("first")
	// Defer runs the statment at the end of the function
	defer fmt.Println("second")
	fmt.Println("third")
	fmt.Println("end of fun def1")
}

func def2() {
	// Defer is executed in LIFO order
	// So it's gonna reverse the order here
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("end of fun def2")
}	

func main() {
	def1()
	def2()
	fmt.Println("end of main")
}
