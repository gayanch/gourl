package handler

import (
    "net/http"
    "fmt"

    "github.com/gayanch/gourl/manager"
)

func Api(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.URL)

    r.ParseForm()
    if len(r.Form) != 2 {
        fmt.Fprintf(w, "{response: invalid request}")
    } else {
        api := r.Form["api"][0]
        switch api {
        case "short":
            longurl := r.Form["url"][0]
            shorturl, _ := manager.SaveUrl(longurl)
            fmt.Fprintf(w, "{response: %s}", shorturl)

        case "long":
            shorturl := r.Form["url"][0]
            if longurl, err := manager.ReadUrl(shorturl); err == nil {
                fmt.Fprintf(w, "{response: %s}", longurl)
            } else {
                fmt.Fprintf(w, "{response: 'not found'}")
            }
        default:
            fmt.Fprintf(w, "{response: invalid request}")
        }
    }
}
