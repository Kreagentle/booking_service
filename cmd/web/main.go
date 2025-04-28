package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/Kreqgentle/booking_service/pkg/config"
	"github.com/Kreqgentle/booking_service/pkg/handlers"
	"github.com/Kreqgentle/booking_service/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
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
