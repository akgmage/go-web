package main

import (
	"fmt"
	"net/http"

	"github.com/akgmage/go-web/pkg/handlers"
)

const portNumber = ":8080"
// Func name starts with Uppercase is visible outside package
// lowercase it is not


// main is the main application function
func main() {
	// listens for a request sent by a web browser
	
	http.HandleFunc("/", handlers.Home) // register Home handler
	
	http.HandleFunc("/about", handlers.About) // register About handler
	
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// start server to listen for request
	_ = http.ListenAndServe(portNumber, nil)
}