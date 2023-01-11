package config

import "html/template"

// AppConfig holds the application config
// immediately available to every other part of the application
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
}