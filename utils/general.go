package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateToken(Username string) string {
	hasher := md5.New()
	hasher.Write([]byte(Username))
	return hex.EncodeToString(hasher.Sum(nil))
}
