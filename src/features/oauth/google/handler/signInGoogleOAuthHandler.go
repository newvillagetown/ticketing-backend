package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"main/common"
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
// @Success 200 {object} bool
// @Failure 400 {object} common.ResError
// @Failure 500 {object} common.ResError
// @Tags auth
func (s *SignInGoogleOAuthHandler) SignInGoogle(c echo.Context) error {
	//콜백 url을 호출
	fmt.Println("signIn")
	fmt.Println(common.OAuthConf.AuthCodeURL("state"))
	c.Redirect(http.StatusMovedPermanently, common.OAuthConf.AuthCodeURL("state"))
	return c.JSON(http.StatusOK, true)
}
