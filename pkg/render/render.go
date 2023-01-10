package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// renderTemplate takes a response writer , name of template, parse it and
// write it to browser window

// When someone visit site and look at our page, this function RenderTemplate2
// every single time reads from disk (layout file) and actual template we want to render
// Read from disk, then parsed and stored in variable which is not every effecient reading
// from disk on every single request.
func RenderTemplate2(w http.ResponseWriter, tmpl string) {
	// load file from root of application
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl") // include base layout template
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

// renderTemplate takes a response writer , name of template, parse it and
// write it to browser window
// Build a Simple template cache
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		// create the template
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache
	tc[t] = tmpl
	return nil
}
