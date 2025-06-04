package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kreqgentle/booking_service/internal/config"
	"github.com/Kreqgentle/booking_service/internal/models"
	"github.com/Kreqgentle/booking_service/internal/render"
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

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (rep *Repository) BookPostJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Availiable!",
	}

	respJson, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJson)
}

func (rep *Repository) MakeBooking(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-booking.page.tmpl", &models.TmpltData{})
}

func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TmpltData{})
}
