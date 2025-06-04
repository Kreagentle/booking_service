package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/Kreqgentle/booking_service/internal/config"
	"github.com/Kreqgentle/booking_service/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	/*mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Rpstr.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Rpstr.About))*/
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Rpstr.Home)
	mux.Get("/about", handlers.Rpstr.About)
	mux.Get("/contact", handlers.Rpstr.Contact)
	mux.Get("/room1", handlers.Rpstr.Room1)
	mux.Get("/room2", handlers.Rpstr.Room2)
	mux.Get("/room3", handlers.Rpstr.Room3)

	mux.Get("/book", handlers.Rpstr.Book)
	mux.Post("/book", handlers.Rpstr.BookPost)
	mux.Post("/book-json", handlers.Rpstr.BookPostJson)
	mux.Get("/make-booking", handlers.Rpstr.MakeBooking)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
