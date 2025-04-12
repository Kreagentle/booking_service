package main

import (
	"net/http"

	"github.com/Kreqgentle/booking_service/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "Hello World")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "About")
}
