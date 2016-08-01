# GOURL
 Simple URL Shortener Web Application. Demo is live at http://gourl.herokuapp.com

# How to get
* `go get github.com/gayanch/gourl`
* Build
* Copy 'static' & 'template' folders to 'bin' directory
* Run

# Usage
 * Goto `src/github.com/gayanch/gourl/handler.go` and change `SITE_ADDRESS` according to your domain name
 * Customize front-end by changing `bin/template/home.html`
 * Build & Run

### Using API
 * Generate short URL
  * `<host>/api?api=short&url=www.google.lk` will return JSON output `{response: <4 letter shortcode>}`

 * Query short URL
  * `<host>/api?api=long&url=<4 letter shortcode>` will return JSON output `{response: <long url>}`


# Notes
* Normally app runs on localhost:8080 (set environment variable to override `PORT=<port_no>`)
* Short URLs are stored in plain-text files, no databases

# To-Do
- [x] Check whether particular long URL is already shortened before generating a new short URL
- [x] Drop Linux dependability
- [x] API functions
- [ ] More customization options without editing source code
- [x] Improve front-end's look & feel
- [ ] many more...
