package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterGoogleOAuthHandler(e *echo.Group) {
	handler := NewGoogleOAuthHandler()
	e.GET("/google/signin", handler.SignInGoogleOAuthHandler.SignInGoogle)
	e.GET("/google/signin/callback", handler.CallbackGoogleOAuthHandler.GoogleSignInCallback)
}
