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

func def3(){
	// This one is tricky, as defer is evaluated at the place in code it is
	// So "start" is going to be printed
	a := "start"
	defer println(a)
	a = "end"
}

func main() {
	def1()
	def2()
	def3()
	fmt.Println("end of main")
}
