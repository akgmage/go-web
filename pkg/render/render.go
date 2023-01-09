package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate takes a response writer , name of template, parse it and
// write it to browser window
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// load file from root of application
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

}