package handlers

import (
	"net/http"

	"github.com/bashbeebo/gommerce/pkg/config"
	"github.com/bashbeebo/gommerce/pkg/render"
)
var Repo *Repository
type Repository struct{
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
