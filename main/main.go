package main

import (
	"net/http"
	"os"

	"github.com/gayanch/gourl/handler"
	"github.com/gayanch/gourl/config"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/api", handler.Api)
	http.HandleFunc("/", handler.Home)

	if port := os.Getenv("PORT"); port == "" {
		if port, ok := config.ConfigValues["port"]; ok {
			http.ListenAndServe(":" + port, nil)
		} else {
			http.ListenAndServe(":8080", nil)
		}
	} else {
		//environment variable takes highest priority
		http.ListenAndServe(":" + port, nil)
	}
}
