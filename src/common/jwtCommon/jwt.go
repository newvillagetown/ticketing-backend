package jwtCommon

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"main/common/envCommon"
	"time"
)

type jwtCustomClaims struct {
	Email      string `json:"email"`
	createTime int64  `json:"createTime"`
	jwt.StandardClaims
}

var AccessToknenSecretKey []byte
var RefreshTokenSecretKey []byte
var JwtConfig middleware.JWTConfig

const (
	AccessTokenExpiredTime  = 1
	RefreshTokenExpiredTime = 24 * 7
)

func InitJwt() error {
	secret := "secret"
	AccessToknenSecretKey = []byte(secret)
	RefreshTokenSecretKey = []byte(secret)
	return nil
}

func GenerateToken(email string) (string, string, error) {
	now := time.Now()
	accessToken, err := GenerateAccessToken(email, now)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenerateRefreshToken(email, now)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func GenerateAccessToken(email string, now time.Time) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		email,
		envCommon.TimeToEpocMillis(now),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * AccessTokenExpiredTime).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString(AccessToknenSecretKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func GenerateRefreshToken(email string, now time.Time) (string, error) {
	claims := &jwtCustomClaims{
		email,
		envCommon.TimeToEpocMillis(now),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * RefreshTokenExpiredTime).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	refreshToken, err := token.SignedString(RefreshTokenSecretKey)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
