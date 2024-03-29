package helper

import (
	"Cloud/core/define"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	// id
	// identity
	// name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKer))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
