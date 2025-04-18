package config

import "html/template"

// AddConfig holds the app config
type AddConfig struct {
	TemplateCache map[string]*template.Template
}
