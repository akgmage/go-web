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
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl") // include base layout template
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

}

tc := make(map[string]template.Template)

func RenderTempalte2(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		// create the template
	} else {
		log.Println("Using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
}
