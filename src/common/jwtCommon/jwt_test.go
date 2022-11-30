package jwtCommon

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateAccessToken(t *testing.T) {
	t.Run("generate accessToken test", func(t *testing.T) {
		now := time.Now()
		email := "test@gmail.co.kr"
		userID := "632acxzct34s"
		InitJwt()
		accessToken, _ := GenerateAccessToken(email, now, userID)
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}
			return AccessToknenSecretKey, nil
		}
		got, _ := jwt.Parse(accessToken, keyFunc)

		want := jwt.MapClaims{}
		want["email"] = email
		want["exp"] = float64(time.Now().Add(time.Hour * AccessTokenExpiredTime).Unix())
		want["userID"] = userID
		assert.Equal(t, want, got.Claims)
	})
}

func TestGenerateRefreshToken(t *testing.T) {
	t.Run("generate accessToken test", func(t *testing.T) {
		now := time.Now()
		email := "test@gmail.co.kr"
		userID := "632acxaczc4s"
		InitJwt()
		refreshToken, _ := GenerateRefreshToken(email, now, userID)
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}
			return RefreshTokenSecretKey, nil
		}
		got, _ := jwt.Parse(refreshToken, keyFunc)

		want := jwt.MapClaims{}
		want["email"] = email
		want["exp"] = float64(time.Now().Add(time.Hour * RefreshTokenExpiredTime).Unix())
		want["userID"] = userID
		assert.Equal(t, want, got.Claims)
	})
}
