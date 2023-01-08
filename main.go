package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"
// Func name starts with Uppercase is visible outside package
// lowercase it is not

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

}






// main is the main application function
func main() {
	// listens for a request sent by a web browser
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// start server to listen for request
	_ = http.ListenAndServe(portNumber, nil)
}