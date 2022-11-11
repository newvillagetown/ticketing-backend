package handler

import (
	"github.com/labstack/echo/v4"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
	_interface "main/features/oauth/google/usecase/interface"
	"net/http"
)

type SignInGoogleOAuthHandler struct {
	UseCase _interface.ISignInGoogleOAuthUseCase
}

func NewSignInGoogleOAuthHandler() *SignInGoogleOAuthHandler {
	return &SignInGoogleOAuthHandler{UseCase: usecase.NewSignInGoogleOAuthUseCase(repository.NewSignInGoogleOAuthRepository())}
}

// GoogleSignin
// @Router /v0.1/auth/google/signin [get]
// @Summary google 로그인
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} data.ResJustOk
// @Failure 400 {object} data.ErrResponse
// @Failure 500 {object} data.ErrResponse
// @Tags auth
func (s *SignInGoogleOAuthHandler) SignInGoogle(c echo.Context) error {
	err := s.UseCase.SignInGoogle()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "")
}
