package handler

import "github.com/labstack/echo/v4"

func RegisterGoogleOAuthHandler(e *echo.Group) {
	handler := NewGoogleOAuthHandler()
	api := e.Group("/auth/google")
	api.GET("/signin", handler.SignInGoogleOAuthHandler.SignInGoogle)
	api.GET("/signin/callback", handler.CallbackGoogleOAuthHandler.GoogleSignInCallback)
}
