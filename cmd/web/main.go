package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/akgmage/go-web/pkg/config"
	"github.com/akgmage/go-web/pkg/handlers"
	"github.com/akgmage/go-web/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
// Func name starts with Uppercase is visible outside package
// lowercase it is not


// main is the main application function
func main() {

	

	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	// Persist sets whether the session cookie should be persistent or not 
	// (i.e. whether it should be retained after a user closes their browser
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction 

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

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	


	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	
	// start server to listen for request
	err = srv.ListenAndServe()
	log.Fatal(err)
}