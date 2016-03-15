package manager

import (
	"crypto/sha1"
	"encoding/hex"
)

//hashes longurls using sha1, so they can be used as filenames without errors
func UrlToHash(longurl string) string {
	hasher := sha1.New()
	hasher.Write([]byte(longurl))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
