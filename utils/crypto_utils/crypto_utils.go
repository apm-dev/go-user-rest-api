package crypto_utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func GetMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetSHA256(input string) string {
	hash := sha256.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
