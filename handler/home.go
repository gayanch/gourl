package handler

import (
  "net/http"
  "fmt"
)

func Home(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.Method, r.URL)

  if (len(r.URL.Path) == 1) {
    //no short url, serve home
    fmt.Fprintf(w, "Home")
  } else {
    //short url provided, redirect
    short := r.URL.Path[len("/"):]
    fmt.Println(short)
    fmt.Fprintf(w, "Redirect to %s", short)
  }

}
