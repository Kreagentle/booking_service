package main

import (
	"encoding/gob"
	"fmt"
	"github.com/Kreqgentle/booking_service/internal/handlers"
	"github.com/Kreqgentle/booking_service/internal/models"
	"github.com/Kreqgentle/booking_service/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"

	"github.com/Kreqgentle/booking_service/internal/config"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	serve := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		fmt.Println("We have an error %w", err)
		return
	}
	/*err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("We have an error %w", err)
		return
	}*/
}

func run() error {
	gob.Register(models.Reservation{})

	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	cache, err := render.CreateCacheTemplate()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)
	// http.HandleFunc("/", handlers.Rpstr.Home)
	// http.HandleFunc("/about", handlers.Rpstr.About)
	return nil
}
