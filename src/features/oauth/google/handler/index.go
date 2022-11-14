package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterGoogleOAuthHandler(e *echo.Group) {
	handler := NewGoogleOAuthHandler()
	e.GET("/signin", handler.SignInGoogleOAuthHandler.SignInGoogle)
	e.GET("/signin/callback", handler.CallbackGoogleOAuthHandler.GoogleSignInCallback)
}
