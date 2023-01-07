package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Go Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(10, 11)
	fmt.Fprint(w, fmt.Sprintf("This is Go About page and 10 + 11 is %d", sum))
}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	// listens for a request sent by a web browser
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// start server to listen for request
	_ = http.ListenAndServe(":8080", nil)
}