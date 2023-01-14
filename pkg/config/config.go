package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config
// immediately available to every other part of the application
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	InProduction bool
}