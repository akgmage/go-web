package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate takes a response writer , name of template, parse it and
// write it to browser window

// When someone visit site and look at our page, this function RenderTemplate2
// every single time reads from disk (layout file) and actual template we want to render
// Read from disk, then parsed and stored in variable which is not every effecient reading
// from disk on every single request.
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache

	// get the requested template from cache

	// render the template

	// load file from root of application
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl") // include base layout template
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
}



// var tc = make(map[string]*template.Template)

// // renderTemplate takes a response writer , name of template, parse it and
// // write it to browser window
// // Build a Simple template cache
// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]
// 	// check to see if we already have template in cache
// 	if !inMap {
// 		// not in map(cache) create the template
// 		log.Println("Creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		// check for error
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// in map, we have template in map (cache)
// 		log.Println("Using cached template")
// 	}
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	// check for error
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// // createTemplateCache parse all required files for an individual template and puts template into cache
// func createTemplateCache(t string) error {
// 	// parse all required files for an individual template
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	// parse the template (take each entry from slice of strings (templates) and put them in as individual strings)
// 	tmpl, err := template.ParseFiles(templates...)
// 	// check for error
// 	if err != nil {
// 		return err
// 	}
// 	// add template to map (cache)
// 	tc[t] = tmpl
// 	return nil
// }
