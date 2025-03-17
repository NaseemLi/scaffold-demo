package jwtutil

import (
	"errors"
	"scaffold-demo/config"
	"scaffold-demo/utils/logs"
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
		username,
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

// 解析token
func ParseToken(ss string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})

	if err != nil {
		//失败
		logs.Error(nil, "解析Token失败")
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//说明token合法
		return claims, nil
	} else {
		//说明失败token不合法
		logs.Error(nil, "解释Token失败")
		return nil, errors.New("Token不合法:invalid token")
	}
}
