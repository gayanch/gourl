package manager

import (
	"errors"
	"fmt"
	"net/url"
	"os"
)

const (
	//dir for save urls
	URL_DIR      = "shorturls/"
	LONG_URL_DIR = "longurls/"
)

//Saves the given longurl in filesystem and returns the generated shorturl
func SaveUrl(longurl string) (string, error) {
	var (
		shorturl string
		err error
	)

	os.Mkdir(URL_DIR, os.ModePerm)
	os.Mkdir(LONG_URL_DIR, os.ModePerm)

	if longurl, err = FormatUrl(longurl); err != nil {
		return "", err
	}

	//escape url
	var furl *url.URL
	if furl, err = url.Parse(longurl); err == nil {
		longurl = furl.String()
	} else {
		return "", errors.New("Invalid URL")
	}

	//check whether short url for given long url is already generated
	//hash longurls before saving/retrieving to avoid file-system restrictions
	filename := UrlToHash(longurl)

	var longUrlFile *os.File
	if longUrlFile, err = os.Open(LONG_URL_DIR + filename); err == nil {
		//generated shorturl found
		fmt.Fscanf(longUrlFile, "%s", &shorturl)
		longUrlFile.Close()

		return shorturl, nil
	} else {
		//no generated url found, generate new one
		//repeat while new short code is found
		for {
			shorturl = GenerateUrl()
			var urlfile *os.File
			if urlfile, err = os.Open(URL_DIR + shorturl); err != nil {
				//create shorturl file
				urlfile, err = os.Create(URL_DIR + shorturl)
				fmt.Fprintf(urlfile, "%s", longurl)
				urlfile.Close()

				//create longurl file
				longUrlFile, _ = os.Create(LONG_URL_DIR + filename)
				fmt.Fprintf(longUrlFile, "%s", shorturl)
				longUrlFile.Close()

				break
			} else {
				urlfile.Close()
			}
		}
		return shorturl, nil
	}
	return "", errors.New("URL Saving error")
}
