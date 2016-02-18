package main

import (
  "net/http"

  "github.com/gayanch/gourl/handler"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handler.Home)

    http.ListenAndServe(":8080", nil)
}

// import (
//     "github.com/gayanch/gourl/manager"
// )
//
// func main() {
//     manager.SaveUrl("http://www.facebook.com/gayanch")
// }
