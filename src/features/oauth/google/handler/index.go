package handler

import "github.com/labstack/echo/v4"

func RegisterGoogleOAuthHandler(e *echo.Group) {
	handler := NewGoogleOAuthHandler()
	api := e.Group("/auth/google")
	api.POST("/login", handler.SignInGoogleOAuthHandler.SignInGoogle)
	api.POST("/callback", handler.CallbackGoogleOAuthHandler.GoogleSignInCallback)
}
