package handlers

import (
	"fmt"
	"net/http"

	"github.com/Kreqgentle/booking_service/pkg/config"
	"github.com/Kreqgentle/booking_service/pkg/models"
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
	remoteIP := r.RemoteAddr
	rep.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TmpltData{})
}

func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	mpString := map[string]string{}
	mpString["test"] = "About test string"

	remoteIP := rep.App.Session.GetString(r.Context(), "remote_ip")
	mpString["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TmpltData{MpString: mpString})
}

func (rep *Repository) Room1(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "room1.page.tmpl", &models.TmpltData{})
}

func (rep *Repository) Room2(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "room2.page.tmpl", &models.TmpltData{})
}

func (rep *Repository) Room3(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "room3.page.tmpl", &models.TmpltData{})
}

func (rep *Repository) Book(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "book.page.tmpl", &models.TmpltData{})
}

func (rep *Repository) BookPost(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s\n", start, end)))
}

func (rep *Repository) MakeBooking(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-booking.page.tmpl", &models.TmpltData{})
}

func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TmpltData{})
}
