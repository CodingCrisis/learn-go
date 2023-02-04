package main

import (
	"fmt"
	"net/http"
	"log"
)

func pan1(){
	// Regular panic due to error
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}

func pan2(){
	// Custom panic
	fmt.Println("start")
	// panic happens after defered statements
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}

func pan3(){
	// Custom panic
	fmt.Println("start")
	// Defering an anonymous function call - a call, so there are parentheses at the end
	defer func(){
		// recover returns nil if everything is ok (app is not panicing)
		// else it's returning the actual error
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}

// Custom web app
// returns Hello Go, on http://localhost:8080/
// panics when port is blocked, so trying to run a second copy of the app will panic
func pan4(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func main(){
	//pan1()
	//pan2()
	pan3()
	//pan4()
}
