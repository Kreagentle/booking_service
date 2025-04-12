package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmplt string) {
	parsedTmpl, _ := template.ParseFiles("./templates" + tmplt)
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template %w", err)
	}
}
