package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Go Home Page")
}

func main() {
	// listens for a request sent by a web browser
	http.HandleFunc("/", Home)

	// start server to listen for request
	_ = http.ListenAndServe(":8080", nil)
}