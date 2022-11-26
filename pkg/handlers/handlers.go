package handlers

import (
	"github.com/DirksCGM/dirkscgm-website/pkg/render"
	"net/http"
)

// Home is the handler function for a home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}
