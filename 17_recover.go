package main

import(
	"fmt"
	"log"
)

func main(){
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker(){
	fmt.Println("about to panic")
	defer func(){
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// if the error is serious we can re-panic
			// in that case the main function will panic as well
			// without the panic, only this internal function will stop executing
			// main will live on and print the "end" text
			//panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
