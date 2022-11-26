package main

import (
	"fmt"
	"github.com/DirksCGM/dirkscgm-website/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the entrypoint for the web application
func main() {
	// ToDo: Complete the rendering & templateCache

	// templates with their layouts are handled here, rendered via the render func
	http.HandleFunc("/", handlers.Home)

	fmt.Printf("Web server running on: http://localhost%s/", portNumber)

	// start the webserver
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}
