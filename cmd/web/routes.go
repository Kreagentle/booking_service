package main

import (
	"net/http"

	"github.com/Kreqgentle/booking_service/pkg/config"
	"github.com/Kreqgentle/booking_service/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	/*mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Rpstr.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Rpstr.About))*/
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Rpstr.Home)
	mux.Get("/about", handlers.Rpstr.About)

	return mux
}
