package jwtCommon

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var JwtSecretKey []byte
var JwtConfig middleware.JWTConfig

func InitJwt() error {
	secret := "secret"
	JwtSecretKey = []byte(secret)
	return nil
}
