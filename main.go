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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, error := divideValues(100.0, 10.0)
	if error != nil {
		fmt.Fprintf(w, "cannot divide by zero")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {
	// listens for a request sent by a web browser
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// start server to listen for request
	_ = http.ListenAndServe(portNumber, nil)
}