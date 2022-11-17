package jwtCommon

import "github.com/golang-jwt/jwt"

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var JwtSecretKey []byte

func InitJwt() error {
	secret := "secret"
	JwtSecretKey = []byte(secret)
	return nil
}
