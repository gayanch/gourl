package manager

import (
	"errors"
	"regexp"
)

var (
	reHttp *regexp.Regexp = nil
	reFtp *regexp.Regexp = nil
	reUrl *regexp.Regexp = nil
)

func init () {
	reHttp, _ = regexp.Compile(`^http[s]{0,1}://.+\..+`)
	reFtp, _ = regexp.Compile(`^ftp[s]{0,1}://.+\..+`)
	reUrl, _ = regexp.Compile(`.+\..+`)
}

const (
	HTTP_PREFIX  = "http://"
//	HTTPS_PREFIX = "https://"
//	FTP_PREFIX   = "ftp://"
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
		
		case reHttp.MatchString(url), reFtp.MatchString(url):
			returl = url
		
		case reUrl.MatchString(url):
			returl = HTTP_PREFIX + url
			
		default:
			err = errors.New("Invalid URL")
			
//	case len(url) <= len(HTTP_PREFIX):
//		returl = HTTP_PREFIX + url

//	case url[:len(HTTP_PREFIX)] == HTTP_PREFIX,
//		url[:len(HTTPS_PREFIX)] == HTTPS_PREFIX,
//		url[:len(FTP_PREFIX)] == FTP_PREFIX:
//		returl = url

//	case url[:len(HTTP_PREFIX)] != HTTP_PREFIX:
//		returl = HTTP_PREFIX + url
	}

	return returl, err
}
