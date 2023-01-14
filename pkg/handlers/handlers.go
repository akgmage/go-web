package handlers

import (
	"net/http"

	"github.com/akgmage/go-web/pkg/config"
	"github.com/akgmage/go-web/pkg/models"
	"github.com/akgmage/go-web/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository
// Repository is a repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func Newhandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
// (m *Repository) with this receiver we have access to everything inside Repository i.e application config
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic 
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	// send data to the template

	render.RenderTemplate(w, "about.page.tmpl",  &models.TemplateData{
		StringMap: stringMap,
	})
}


