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
	mux.Get("/contact", handlers.Repo.HandleContact)
	mux.Get("/detail", handlers.Repo.HandleNewsDetail)
	mux.Get("/generals-quarters", handlers.Repo.HandleGenerals)
	mux.Get("/majors-suite", handlers.Repo.HandleMajors)
	mux.Get("/search-availability", handlers.Repo.HandleSearchAvailability)
	mux.Post("/search-availability", handlers.Repo.HandlePostSearchAvailability)

	mux.Get("/make-reservation", handlers.Repo.HandleMakeReservation)

	/* TAMBAHKAN DISINI UNTUK MELOAD DATA STATIC: images, css, js dll */
	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
