package main

import(
	"fmt"
	"time"
)

func main(){
	var msg = "Hello"
	// Using an anonymous function
	go func() {
		// Interestingly enoguh go has access to outer scope (closures)
		// This is problematic, because of the dependency
		fmt.Println("extern param", msg)
	}()

	go func(msg string){
		fmt.Println("with param", msg)
	}(msg)
		
	// Usually the go routine is run after the rest of the parent function
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
