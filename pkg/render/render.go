package render

import (
	"bytes"
	"github.com/DirksCGM/dirkscgm-website/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders the template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var templateCache map[string]*template.Template

	// rebuild the templates on every request if in dev mode
	if app.UseCache {
		// get template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// Get template requested from cache
	t, isOk := templateCache[tmpl]
	if !isOk {
		log.Fatal("could not get the template from template cache") // kills the process
	}

	buffer := new(bytes.Buffer)   // has bytes to execute values of the map from the buffer (better error checking)
	err := t.Execute(buffer, nil) // if parsed but fails, a useful indication is given as to why
	if err != nil {
		log.Println(err)
	}

	// Render the template by writing response writer to buffer
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache returns a map of string pointer of Template
// this means we don't need to keep track of templates and layouts
func CreateTemplateCache() (map[string]*template.Template, error) {
	// create an empty template map for the cache called cache
	cache := map[string]*template.Template{}

	// populate the entire cache on load by referring to the filesystem
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return cache, err
	}

	// range through all *.page.gohtml
	for _, page := range pages {
		name := filepath.Base(page)                    // get name of file
		ts, err := template.New(name).ParseFiles(page) // parse file and store it in a template called itself
		if err != nil {
			return cache, err
		}

		// look for layouts in dir
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return cache, err
		}

		// only parse base if there are any
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts // apply template set to the cache
	}

	return cache, nil
}
