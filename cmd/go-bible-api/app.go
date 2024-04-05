package main

import (
	"log"
	"net/http"

    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Router *chi.Mux
	Config *Config
}

func NewApp(config *Config) *App {
	app := &App{
		Router: chi.NewRouter(),
		Config: config,
	}

	app.initialize()

	return app
}

func (app *App) initialize() {
    app.Router.Use(middleware.Logger)

    app.Router.Get("/", Root)
}

func (app *App) Run() {
	log.Printf("Netcalc API Server Started [%s]\n", app.Config.HttpPort)
	log.Println(http.ListenAndServe(":"+app.Config.HttpPort, app.Router))
}
