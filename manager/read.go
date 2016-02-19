package manager

import (
    "os"
    "fmt"
)

//returns the longurl of given shorturl, err != nil if given shorturl is not found
func ReadUrl(shorturl string) (string, error) {
    var (
        urlfile *os.File
        longurl string
        err error
    )

    if urlfile, err = os.Open(URL_DIR + shorturl); err == nil {
        fmt.Fscanf(urlfile, "%s", &longurl)
        urlfile.Close()
    }

    return longurl, err
}
