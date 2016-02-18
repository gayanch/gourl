package manager

import (
    "os"
    "fmt"
)

const (
    //max length of a short url
    MAX_URL_LEN = 4
)

//generates a randrom string with length of MAX_URL_LEN
//To-Do:
//use go lang's rand package instead of uding linux specific method
func GenerateUrl() string {
    var (
        char rune
        url string
    )

    //Linux (may be unix/Mac) only
    file, _ := os.Open("/dev/urandom")
    defer file.Close()

    for i := 0; i < MAX_URL_LEN; {
        fmt.Fscanf(file, "%c", &char)

        //check whether char is a ascii value
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
            url += string(char)
            i++
        }
    }
    return url
}
