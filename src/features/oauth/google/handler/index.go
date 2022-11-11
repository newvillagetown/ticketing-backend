package handler

import "github.com/labstack/echo/v4"

func RegisterAddressBookHandler(e *echo.Group) {
	handler := NewGoogleOAuthHandler()
	api := e.Group("/v0.1")
	api.POST("/v0.1", handler.SignInGoogleOAuthHandler.SignInGoogle)
}
