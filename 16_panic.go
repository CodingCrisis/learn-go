package main

import (
	"fmt"
	"net/http"
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
	panic("something bad happened")
	fmt.Println("end")
}

// Custom web app
// returns Hello Go, on http://localhost:8080/
// panics when port is blocked, so trying to run a second copy of the app will panic
func pan3(){
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
}
