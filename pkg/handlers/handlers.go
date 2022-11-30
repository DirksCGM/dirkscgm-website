// Note that we will be using a repository design pattern
// more on it here: https://threedots.tech/post/repository-pattern-in-go/

package handlers

import (
	"github.com/DirksCGM/dirkscgm-website/pkg/config"
	"github.com/DirksCGM/dirkscgm-website/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository is the repo type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repo for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler function for a home page with a
// receiver to link all handlers with the repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}
