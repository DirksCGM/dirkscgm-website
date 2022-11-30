package main

import (
	"fmt"
	"github.com/DirksCGM/dirkscgm-website/pkg/config"
	"github.com/DirksCGM/dirkscgm-website/pkg/handlers"
	"github.com/DirksCGM/dirkscgm-website/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the entrypoint for the web application
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	// pass the app to the repository handler for site-wide access
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// templates with their layouts are handled here, rendered via the render func
	http.HandleFunc("/", handlers.Repo.Home)
	fmt.Printf("Web server running on: http://localhost%s/", portNumber)

	// start the webserver
	_ = http.ListenAndServe(portNumber, nil)
}
