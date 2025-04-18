package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kreqgentle/booking_service/pkg/config"
	"github.com/Kreqgentle/booking_service/pkg/handlers"
	"github.com/Kreqgentle/booking_service/pkg/render"
)

func main() {
	var app config.AddConfig

	cache, err := render.CreateCacheTemplate()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = cache
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("We have an error %w", err)
		return
	}
}
