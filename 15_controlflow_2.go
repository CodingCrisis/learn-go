package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	// Using defer here allows opening and closing a resource next to each other in code
	// but the execution of the close is done at the end of the function
	// using it in loops is tricky, so maybe skip it then
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
