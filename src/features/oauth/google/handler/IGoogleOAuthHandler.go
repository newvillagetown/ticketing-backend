package handler

import (
	"github.com/labstack/echo/v4"
	_interface "main/features/oauth/google/usecase/interface"
)

type ISignInGoogleOAuthHandler interface {
	SignInGoogle(c echo.Context) error
}

type IGoogleOAuthHandler interface {
	NewSignInGoogleOAuthHandler(UseCase _interface.ISignInGoogleOAuthUseCase) *SignInGoogleOAuthHandler
}

type GoogleOAuthHandler struct {
	SignInGoogleOAuthHandler SignInGoogleOAuthHandler
}

func NewGoogleOAuthHandler() *GoogleOAuthHandler {
	return &GoogleOAuthHandler{
		SignInGoogleOAuthHandler: *NewSignInGoogleOAuthHandler(),
	}
}
