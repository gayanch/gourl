package manager

import (
	"math/rand"
	"time"

	"github.com/gayanch/gourl/config"
	"strconv"
)

const (
	//string used to generate short codes
	GEN_STRING = "0123456789" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz"
)

//generates a randrom string with length of MAX_URL_LEN
func GenerateUrl() string {
	var (
		url string
		MAX_URL_LEN int
	)

	//getting maximum short code length from configuration values
	if value, ok := config.ConfigValues["url_len"]; ok {
		MAX_URL_LEN, _ = strconv.Atoi(value)
	} else {
		//fallback value if configuration can't be readed
		MAX_URL_LEN = 4
	}

	seed := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(seed)

	for i := 0; i < MAX_URL_LEN; i++ {
		url += string(GEN_STRING[randGen.Intn(len(GEN_STRING))])
	}
	return url
}
