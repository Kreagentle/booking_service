package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("We have an error %w", err)
		return
	}
}
