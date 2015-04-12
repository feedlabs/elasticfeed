package common

import (
	"crypto/md5"
	"encoding/hex"

	auth "github.com/abbot/go-http-auth"
)

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetCrypt(password string) string {
	md5 := string(auth.MD5Crypt([]byte(password), []byte(""), []byte("$$")))
	return md5
}
