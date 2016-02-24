package main

import (
	"net/http"
	"os"

	"github.com/gayanch/gourl/handler"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/api", handler.Api)
	http.HandleFunc("/", handler.Home)

	if port := os.Getenv("PORT"); port == "" {
		http.ListenAndServe(":8080", nil)
	} else {
		http.ListenAndServe(":"+port, nil)
	}
}
