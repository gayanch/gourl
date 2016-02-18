package manager

import (
    //"os"
    //"fmt"
    "math/rand"
    "time"
)

const (
    //max length of a short url
    MAX_URL_LEN = 4

    GEN_STRING = "0123456789" +
                    "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
                    "abcdefghijklmnopqrstuvwxyz"
)

//generates a randrom string with length of MAX_URL_LEN
func GenerateUrl() string {
    var (
        //char rune
        url string
    )

    //Linux (may be unix/Mac) only
    // file, _ := os.Open("/dev/urandom")
    // defer file.Close()
    seed := rand.NewSource(time.Now().UnixNano())
    randGen := rand.New(seed)

    for i := 0; i < MAX_URL_LEN; i++ {
        url += string(GEN_STRING[ randGen.Intn( len(GEN_STRING) ) ])

        // fmt.Fscanf(file, "%c", &char)
        //
        // //check whether char is a ascii value
        // if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
        //     url += string(char)
        //     i++
        // }
    }
    return url
}
