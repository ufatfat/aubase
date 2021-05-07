package util

import (
	"aubase/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenToken (id uint32, name string) (token string, err error) {
	claims := jwt.MapClaims{
		"id": id,
		"name":     name,
		"timestamp": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(config.TOKEN_SECRET_KEY))
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
		return []byte(config.TOKEN_SECRET_KEY), err
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		claims = t.Claims.(jwt.MapClaims)
	}
	return
}

func ValidateToken (token string) (ok bool) {
	claims, err := getClaimsFromToken(token)
	if err != nil {
		return
	}
	ts := int64(claims.(jwt.MapClaims)["timestamp"].(float64))
	now := time.Now().Unix()
	if ts + config.TOKEN_EXPIRE_TIME < now {
		return
	}
	return true
}

func UpdateToken (token string) (updatedToken string) {
	claims, _ := getClaimsFromToken(token)
	ts := int64(claims.(jwt.MapClaims)["timestamp"].(float64))
	now := time.Now().Unix()
	if ts + config.TOKEN_EXPIRE_TIME < now || ts + config.TOKEN_DYING_TIME > now {
		return
	}
	updatedToken, _ = GenToken(uint32(claims.(jwt.MapClaims)["id"].(float64)), claims.(jwt.MapClaims)["name"].(string))
	return
}