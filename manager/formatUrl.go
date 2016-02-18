package manager

const (
    HTTP_PREFIX = "http://"
    HTTPS_PREFIX = "https://"
)

//appends http:// prefix to given longurl if not present
//http:// prefix is required to perform redurect correctly
func FormatUrl(url string) string {
    switch {
        case len(url) < len(HTTP_PREFIX):
            url = HTTP_PREFIX + url

        case url[ :len(HTTPS_PREFIX)] == HTTPS_PREFIX:
            //do nothing

        case url[ :len(HTTP_PREFIX)] != HTTP_PREFIX:
            url = HTTP_PREFIX + url
    }

    return url
}
