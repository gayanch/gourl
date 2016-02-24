package manager

import (
	"math/rand"
	"time"
)

const (
	//max length of a short url
	MAX_URL_LEN = 4

	//string used to generate short codes
	GEN_STRING = "0123456789" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz"
)

//generates a randrom string with length of MAX_URL_LEN
func GenerateUrl() string {
	var (
		url string
	)

	seed := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(seed)

	for i := 0; i < MAX_URL_LEN; i++ {
		url += string(GEN_STRING[randGen.Intn(len(GEN_STRING))])
	}
	return url
}
