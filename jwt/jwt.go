package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(claims MyClaims, key string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(key))
}

func ParseTokenString(tokenString string, key string) (Claims jwt.MapClaims, er error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("请求令牌无效")
	}

	return token.Claims.(jwt.MapClaims), nil
}
