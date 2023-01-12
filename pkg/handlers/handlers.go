package handlers

import (
	"net/http"

	"github.com/akgmage/go-web/pkg/config"
	"github.com/akgmage/go-web/pkg/render"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap 	map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}

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
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic 

	// send data to the template

	render.RenderTemplate(w, "about.page.tmpl")
}


