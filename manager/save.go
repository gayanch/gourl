package manager

import (
    "os"
    "fmt"
    "strings"
)

const (
    //dir for save urls
    URL_DIR = "shorturls/"
    LONG_URL_DIR = "longurls/"
)

func SaveUrl(longurl string) (shorturl string, err error) {
    os.Mkdir(URL_DIR, os.ModePerm)
    os.Mkdir(LONG_URL_DIR, os.ModePerm)

    longurl = FormatUrl(longurl)

    //check whether short url for given long url is already generated
    replacer := strings.NewReplacer("/", "$")
    filename := replacer.Replace(longurl)

    if longUrlFile, err := os.Open(LONG_URL_DIR + filename); err == nil {
        //generated shorturl found
        fmt.Fscanf(longUrlFile, "%s", &shorturl)
        longUrlFile.Close()
    } else {
        //no generated url found, generate new one
        for {
            shorturl = GenerateUrl()
            if urlfile, err := os.Open(URL_DIR + shorturl); err != nil {
                //create shorturl file
                urlfile, err = os.Create(URL_DIR + shorturl)
                fmt.Fprintf(urlfile, "%s", longurl)
                urlfile.Close()

                //create longurl file
                longUrlFile, _ = os.Create(LONG_URL_DIR + filename)
                fmt.Fprintf(longUrlFile, "%s", shorturl)

                break
            }
        }
    }
    return
}
