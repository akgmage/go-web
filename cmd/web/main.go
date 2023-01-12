package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akgmage/go-web/pkg/config"
	"github.com/akgmage/go-web/pkg/handlers"
	"github.com/akgmage/go-web/pkg/render"
)

const portNumber = ":8080"
// Func name starts with Uppercase is visible outside package
// lowercase it is not


// main is the main application function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.Newhandlers(repo)

	// give render package access to app config
	render.NewTemplates(&app)

	// listens for a request sent by a web browser
	
	// http.HandleFunc("/", handlers.Repo.Home) // register Home handler
	
	// http.HandleFunc("/about", handlers.Repo.About) // register About handler
	
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// start server to listen for request
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}