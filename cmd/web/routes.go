package main

import (
	"net/http"

	"github.com/sakurtek/goserver/bookingremyconcept/pkg/config"
	"github.com/sakurtek/goserver/bookingremyconcept/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HandleHome)
	mux.Get("/about", handlers.Repo.HandleAbout)
	mux.Get("/detail", handlers.Repo.HandleNewsDetail)

	return mux
}
