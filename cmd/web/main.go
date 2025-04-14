package main

import (
	"fmt"
	"net/http"

	"github.com/Kreqgentle/booking_service/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("We have an error %w", err)
		return
	}
}
