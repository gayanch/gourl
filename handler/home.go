package handler

import (
    "net/http"
    "html/template"
    "fmt"

    "github.com/gayanch/gourl/manager"
)

const (
    SITE_ADDRESS = "test.com"
)

func Home(w http.ResponseWriter, r *http.Request) {
      fmt.Println(r.Method, r.URL)
      if r.Method == "GET" {
          if len(r.URL.Path) == 1 {
            //no short url, serve home
            t, _ := template.ParseFiles("template/home.html")
            t.ExecuteTemplate(w, t.Name(), nil)
          } else {
            //short url provided, redirect
            shorturl := r.URL.Path[len("/"):]
            if longurl, err := manager.ReadUrl(shorturl); err == nil {
                http.Redirect(w, r, longurl, 303)
            } else {
                fmt.Fprintf(w, "Invalid short url")
            }

          }
      } else {
          r.ParseForm()
          longurl := r.Form["longurl"][0]

          if shorturl, err := manager.SaveUrl(longurl); err == nil {
              fmt.Fprintf(w, "%s/%s", SITE_ADDRESS, shorturl)
          } else {
              fmt.Fprintf(w, "Error generating url")
          }
      }

}
