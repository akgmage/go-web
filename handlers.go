package main

import (
	"fmt"
	"net/http"
	"text/template"
)

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
