package main

import(
	"fmt"
)

func main(){
	// Spin of a green thread and run the function
	go sayHello()
}

func sayHello(){
	fmt.Println("Hello")
}
