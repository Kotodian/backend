package utils

import (
	"encoding/hex"
	"crypto/md5"
)

func Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


