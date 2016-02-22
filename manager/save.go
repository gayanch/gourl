package manager

import (
    "os"
    "fmt"
    "net/url"
)

const (
    //dir for save urls
    URL_DIR = "shorturls/"
    LONG_URL_DIR = "longurls/"
)

//Saves the given longurl in filesystem and returns the generated shorturl
func SaveUrl(longurl string) (shorturl string, err error) {
    os.Mkdir(URL_DIR, os.ModePerm)
    os.Mkdir(LONG_URL_DIR, os.ModePerm)

    longurl, err = FormatUrl(longurl)

    //escape url
    if url, err := url.Parse(longurl); err == nil {
        longurl = url.String()
    }

    //check whether short url for given long url is already generated
    //hash longurls before saving/retrieving to avoid file-system restrictions
    filename := UrlToHash(longurl)

    if longUrlFile, err := os.Open(LONG_URL_DIR + filename); err == nil {
        //generated shorturl found
        fmt.Fscanf(longUrlFile, "%s", &shorturl)
        longUrlFile.Close()
    } else {
        //no generated url found, generate new one
        //repeat while new short code is found
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
            } else {
                urlfile.Close()
            }
        }
    }
    return
}
