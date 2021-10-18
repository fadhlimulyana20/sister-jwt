package helper

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct{}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	UserId    string
}

var signKey *rsa.PrivateKey

func (j *JWT) CreateToken(user_id string, expHours int, secret string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod(jwt.SigningMethodHS256.Name))

	t.Claims = &CustomClaimsExample{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expHours)).Unix(),
		},
		TokenType: "level1",
		UserId:    user_id,
	}

	return t.SignedString([]byte(secret))
}

func (j *JWT) Parse(tokenString string, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(*CustomClaimsExample)
	return claims.UserId, nil
}
