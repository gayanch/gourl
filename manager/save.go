package manager

import (
    "os"
    "fmt"
)

const (
    //dir for save urls
    URL_DIR = "urls/"
)

func SaveUrl(longurl string) (shorturl string, err error) {
    os.Mkdir(URL_DIR, os.ModePerm)

    for {
        shorturl = GenerateUrl()
        if urlfile, err := os.Open(URL_DIR + shorturl); err != nil {
            urlfile, err = os.Create(URL_DIR + shorturl)
            fmt.Fprintf(urlfile, "%s", longurl)
            urlfile.Close()
            break
        }
    }
    return
}
