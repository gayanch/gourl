package manager

import (
    "os"
    "fmt"
)

const (
    //max length of a short url
    MAX_URL_LEN = 4
)

func GenerateUrl() string {
    var (
        char rune
        url string
    )
    
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
