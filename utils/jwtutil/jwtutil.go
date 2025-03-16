package jwtutil

import (
	"scaffold-demo/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

// 1.自定义声明类型
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 2. 封装生成token的函数
func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "sb",
			Subject:   "lizeyu",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	return ss, err
}
