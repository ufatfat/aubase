package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func PasswordEncrypt (password string) (encrypted string) {
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	encrypted = hex.EncodeToString(cipherStr)
	return
}
func GetIndexOfElem (elems *[]uint32, val uint32) (idx int) {
	for k, v := range *elems {
		fmt.Println("k:", k, ", v:", v)
		if v == val {
			idx = k
			return
		}
	}
	return -1
}