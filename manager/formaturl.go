package manager

import (
	"errors"
)

const (
	HTTP_PREFIX  = "http://"
	HTTPS_PREFIX = "https://"
	FTP_PREFIX   = "ftp://"
)

//appends http:// prefix to given longurl if not present
//http:// prefix is required to perform redirect correctly
func FormatUrl(url string) (string, error) {

	var (
		returl string = ""
		err    error  = nil
	)

	switch {
	case len(url) == 0:
		//handle empty url error here or in client-side
		err = errors.New("Invalid URL")

	case len(url) <= len(HTTP_PREFIX):
		returl = HTTP_PREFIX + url

	case url[:len(HTTP_PREFIX)] == HTTP_PREFIX,
		url[:len(HTTPS_PREFIX)] == HTTPS_PREFIX,
		url[:len(FTP_PREFIX)] == FTP_PREFIX:
		returl = url

	case url[:len(HTTP_PREFIX)] != HTTP_PREFIX:
		returl = HTTP_PREFIX + url
	}

	return returl, err
}
