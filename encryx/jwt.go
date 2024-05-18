package encryx

import (
	"errors"
	"fmt"
	"github.com/fuckqqcom/pkg/constantx"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJwt(secret string, payloads map[string]any, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}

func ParseJwt(token, secret string) (userId any, err error) {
	payload := jwt.MapClaims{}
	claims, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Header)
		return []byte(secret), nil
	})
	if err != nil {
		return
	}
	if claims.Valid {
		return nil, errors.New("token is valid")
	}
	return payload[constantx.UserId], nil
}
