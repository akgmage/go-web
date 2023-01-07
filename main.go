package main

import (
	"fmt"
	"net/http"
)

func main() {
	// listens for a request sent by a web browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(fmt.Sprintf("Number of bytes written: %d", n))
	})

	// start server to listen for request
	_ = http.ListenAndServe(":8080", nil)
}