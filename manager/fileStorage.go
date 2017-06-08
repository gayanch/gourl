package manager

import (
	"os"
	"net/url"
	"fmt"
	"errors"
)

type fileStorage struct {
	shortUrlDir string
	longUrlDir string
}

func GetStorage(shortUrlDir, longUrlDir string) fileStorage {
	return fileStorage{shortUrlDir, longUrlDir}
}

func (f fileStorage) SaveUrl(longurl string) (string, error) {
	var (
		shorturl string
		err error
	)

	os.Mkdir(f.shortUrlDir, os.ModePerm)
	os.Mkdir(f.longUrlDir, os.ModePerm)

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
	if longUrlFile, err = os.Open(f.longUrlDir + filename); err == nil {
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
			if urlfile, err = os.Open(f.shortUrlDir + shorturl); err != nil {
				//create shorturl file
				urlfile, err = os.Create(f.shortUrlDir + shorturl)
				fmt.Fprintf(urlfile, "%s", longurl)
				urlfile.Close()

				//create longurl file
				longUrlFile, _ = os.Create(f.longUrlDir + filename)
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

func (f fileStorage) ReadUrl(shorturl string) (string, error) {
	var (
		urlfile *os.File
		longurl string
		err     error
	)

	if urlfile, err = os.Open(f.shortUrlDir + shorturl); err == nil {
		fmt.Fscanf(urlfile, "%s", &longurl)
		urlfile.Close()
	}

	return longurl, err
}
