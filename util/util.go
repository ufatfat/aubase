package util

import (
	"crypto/md5"
	"encoding/hex"
)

func PasswordEncrypt (password string) (encrypted string) {
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	encrypted = hex.EncodeToString(cipherStr)
	return
}
