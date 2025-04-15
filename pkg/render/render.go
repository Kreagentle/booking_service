package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmplt string) {
	parsedTmpl, _ := template.ParseFiles("./templates"+tmplt, "./templates/base.layout.tmpl")
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template %w", err)
	}
}
