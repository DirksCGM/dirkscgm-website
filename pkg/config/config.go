// Package config: Instead of making use of a .config or global variables
// all the apps configurations are handled in the AppConfig struct
package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool // can be set to load from disk every time or use cache
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
