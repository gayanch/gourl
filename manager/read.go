package manager

import (
    "os"
    "fmt"
)

func ReadUrl(shorturl string) (longurl string, err error) {
    if urlfile, err := os.Open(URL_DIR + shorturl); err == nil {
        fmt.Fscanf(urlfile, "%s", &longurl)
        urlfile.Close()
    } else {
        fmt.Println(err)
    }
    return
}
