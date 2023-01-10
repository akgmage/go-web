package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// renderTemplate takes a response writer , name of template, parse it and
// write it to browser window

// Advantage over old implementation is we no longer have top keep track
// of how many files are in templates directory and how many are using particular extension.page.tmpl or layout.page.tmpl
// Reading all of required files from disk, parsing them, putting them in a map
// pulling the value out of the map and then rendering it 
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get the requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// hold bytes, try to execute the value we got from map and write it out for error checking
	buff := new(bytes.Buffer) 
	
	err = t.Execute(buff, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template by writing to  writer
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page) // get filename
		ts, err := template.New(name).ParseFiles(page) // parse the file and store template 
		if err != nil { // error check
			return myCache, err
		}
		// look for any layouts that exists in that directory
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil { // error check
			return myCache, err
		}
		// deal with templates if there are matches
		if len(matches) > 0 {
			// ParseGlob parses the template definitions in the files 
			// identified by the pattern and associates the resulting 
			// templates with t in our case ts
			// some of the files that end in page.tmpl might require layout files down here
			// so parse those and add them to template set
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil { // error check
				return myCache, err
			}
		}
		myCache[name] = ts
	}	

	return myCache, nil
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
