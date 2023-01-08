package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"
// Func name starts with Uppercase is visible outside package
// lowercase it is not

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// renderTemplate takes a response writer , name of template, parse it and
// write it to browser window
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// load file from root of application
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

}





// main is the main application function
func main() {
	// listens for a request sent by a web browser
	
	http.HandleFunc("/", Home) // register Home handler
	
	http.HandleFunc("/about", About) // register About handler
	
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// start server to listen for request
	_ = http.ListenAndServe(portNumber, nil)
}