package handlers

import (
	"net/http"

	"github.com/bashbeebo/gommerce/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
