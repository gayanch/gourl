package manager

const (
    HTTP_PREFIX = "http://"
    HTTPS_PREFIX = "https://"
)

//appends http:// prefix to given longurl if not present
//http:// prefix is required to perform redurect correctly
func FormatUrl(url string) (string, error) {

    var (
        returl string = ""
        err error = nil
    )

    switch {
        case len(url) == 0:
            //handle empty url error here or in client-side

        case len(url) < len(HTTP_PREFIX):
            returl = HTTP_PREFIX + url

        case url[ :len(HTTPS_PREFIX)] == HTTPS_PREFIX:
            returl = url

        case url[ : len(HTTP_PREFIX)] == HTTP_PREFIX:
            returl = url

        case url[ :len(HTTP_PREFIX)] != HTTP_PREFIX:
            returl = HTTP_PREFIX + url
    }

    return returl, err
}
