package main

import (
	"fmt"
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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("We have an error %w", err)
		return
	}
}
