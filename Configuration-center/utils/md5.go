package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(str []byte) string {
	h := md5.New()
	return hex.EncodeToString(h.Sum(nil))
}
