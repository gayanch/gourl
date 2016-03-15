package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gayanch/gourl/manager"
)

const (
	INVALID_REQUEST = "{response: null}"
)

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL, time.Now())

	r.ParseForm()
	if len(r.Form) != 2 {
		fmt.Fprintf(w, INVALID_REQUEST)
	} else {
		api := r.Form["api"][0]
		switch api {
		case "short":
			longurl := r.Form["url"][0]
			shorturl, _ := manager.SaveUrl(longurl)
			fmt.Fprintf(w, "{response: '%s'}", shorturl)

		case "long":
			shorturl := r.Form["url"][0]
			if longurl, err := manager.ReadUrl(shorturl); err == nil {
				fmt.Fprintf(w, "{response: '%s'}", longurl)
			} else {
				fmt.Fprintf(w, INVALID_REQUEST)
			}
		default:
			fmt.Fprintf(w, INVALID_REQUEST)
		}
	}
}
