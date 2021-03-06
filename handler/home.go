package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gayanch/gourl/manager"
	"github.com/gayanch/gourl/analytics"
)

const (
	//this is the domain name which this app is hosted, modify it according to domain
	//this will not affect the functionalties of the app, just used to show users shorurls
	SITE_ADDRESS = "gourl.herokuapp.com"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL, time.Now())

	//GET resuest, serve homepage
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

				//update analytics
				analytics.Update(shorturl)
			} else {
				//given shorturl is not found, redirect to homepage
				http.Redirect(w, r, "/", 303)
			}
		}

	} else {
		//POST request, generate shorturl from given longurl and return
		r.ParseForm()
		longurl := r.Form["longurl"][0]

		if shorturl, err := manager.SaveUrl(longurl); err == nil {
			fmt.Fprintf(w, "<a href='/%s'>%s</a>", shorturl, shorturl)

		} else {
			fmt.Fprintf(w, "Error generating URL")
		}
	}
}
