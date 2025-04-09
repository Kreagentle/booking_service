package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println("We have an error %w", err)
	}

	fmt.Printf("Number of bytes: %d\n", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println("We have an error %w", err)
	}

	fmt.Printf("Number of bytes: %d\n", n)
}

func RenderTemplate(w http.ResponseWriter, tmplt string) {
	parsedTmpl, _ := template.ParseFiles("./templates" + tmplt)
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template %w", err)
	}
}
