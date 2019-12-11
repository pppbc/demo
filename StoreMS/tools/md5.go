package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}
