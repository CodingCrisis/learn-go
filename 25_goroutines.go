package main

import(
	"fmt"
	"time"
)

func main(){
	// Spin of a green thread and run the function
	go sayHello()
	
	// Wait a bit, just that the main function does not exit before the sayHello one finishes
	time.Sleep(100 * time.Millisecond)
}

func sayHello(){
	fmt.Println("Hello")
}
