package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bashbeebo/gommerce/pkg/config"
	"github.com/bashbeebo/gommerce/pkg/handlers"
	"github.com/bashbeebo/gommerce/pkg/render"
)

const port string = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	fmt.Printf("Starting Application on port %s", port)
	_ = http.ListenAndServe(port, nil)

}
