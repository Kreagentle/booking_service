package handlers

import (
	"net/http"

	"github.com/Kreqgentle/booking_service/pkg/config"
	"github.com/Kreqgentle/booking_service/pkg/render"
)

var Rpstr *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandler(r *Repository) {
	Rpstr = r
}

func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "Hello World")
}

func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "About")
}
