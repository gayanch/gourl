package handler

import (
    "fmt"
    "net/http"

    "github.com/gayanch/gourl/analytics"
)

func Analytics(w http.ResponseWriter, r *http.Request) {
    if len(r.URL.Path) == len("/analytics/") {
        fmt.Fprintf(w, "No Url Specified")
    } else {
        shorturl := r.URL.Path[len("/analytics/"): ]
        fmt.Fprintf(w, "Showing analytics for %s\n%s", shorturl, analytics.GetAnalyticsService().Get(shorturl))
    }
}
