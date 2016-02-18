package manager

const HTTP_PREFIX = "http://"

func FormatUrl(url string) string {
    if len(url) < len(HTTP_PREFIX) || url[: len(HTTP_PREFIX)] != HTTP_PREFIX {
        url = HTTP_PREFIX + url
    }
    
    return url
}
