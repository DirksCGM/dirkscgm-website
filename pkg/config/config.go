// Package config: Instead of making use of a .config or global variables
// all the apps configurations are handled in the AppConfig struct
package config

import "html/template"

// AppConfig holds the application configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
