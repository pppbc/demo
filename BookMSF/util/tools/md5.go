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

func ValidatePasswd(plainpwd, salt, passwd string) bool {
	//log.Println(Md5Encode(plainpwd + salt))
	return Md5Encode(plainpwd+salt) == passwd
}

func MakePasswd(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}
