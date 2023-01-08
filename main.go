package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"
// Func name starts with Uppercase is visible outside package
// lowercase it is not

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Go Home Page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(10, 11)
	_, _ = fmt.Fprint(w, fmt.Sprintf("This is Go About page and 10 + 11 is %d", sum))
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {
	// listens for a request sent by a web browser
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)


	// start server to listen for request
	_ = http.ListenAndServe(portNumber, nil)
}