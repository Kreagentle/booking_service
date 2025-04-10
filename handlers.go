package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Hello World")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "About")
}

func RenderTemplate(w http.ResponseWriter, tmplt string) {
	parsedTmpl, _ := template.ParseFiles("./templates" + tmplt)
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template %w", err)
	}
}
