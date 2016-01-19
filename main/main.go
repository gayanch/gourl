package main

import (
  "net/http"

  "github.com/gayanch/gourl/handler"
)

func main() {
  http.HandleFunc("/", handler.Home)

  http.ListenAndServe(":8080", nil)
}
