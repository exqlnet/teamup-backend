package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"teamup/config"
	"time"
)


func GenerateToken(userId int32, key string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := jwt.MapClaims{}
	expire := config.Cfg.Get("auth", "expire").Int(0)
	claims["exp"] = time.Now().Add(time.Second * time.Duration(expire)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userId"] = userId
	token.Claims = claims
	return token.SignedString([]byte(key))
}


func ParseToken(tokenString string, key string) (int32, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(key), nil
	})
	if err != nil {
		return -1, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return int32(claims["userId"].(float64)), nil
	} else {
		return -1, errors.New("token is not valid")
	}

}