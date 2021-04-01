package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenToken (id uint32, username string) (token string, err error) {
	claims := jwt.MapClaims{
		"id": id,
		"username":     username,
		"timestamp": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte("aubase@2021"))
	return
}

func GetIDFromToken (token string) (id uint32, err error) {
	claims, err := getClaimsFromToken(token)
	if err != nil {
		return
	}
	id = uint32(claims.(jwt.MapClaims)["id"].(float64))
	return
}

func getClaimsFromToken (token string) (claims jwt.Claims, err error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte("aubase@2021"), err
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		claims = t.Claims.(jwt.MapClaims)
	}
	return
}