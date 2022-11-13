package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"main/common/oauth/google"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
	_interface "main/features/oauth/google/usecase/interface"
	"net/http"
)

type CallbackGoogleOAuthHandler struct {
	UseCase _interface.ICallbackGoogleOAuthUseCase
}

func NewCallbackGoogleOAuthHandler() *CallbackGoogleOAuthHandler {
	return &CallbackGoogleOAuthHandler{UseCase: usecase.NewCallbackGoogleOAuthUseCase(repository.NewCallbackGoogleOAuthRepository())}
}

// google signin callback
// @Router /v0.1/auth/google/signin/callback [get]
// @Summary google login callback
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorSystem.ResError
// @Failure 500 {object} errorSystem.ResError
// @Tags auth
func (cc *CallbackGoogleOAuthHandler) GoogleSignInCallback(c echo.Context) error {
	//인증서버에 액세스 토큰 요청
	token, err := google.OAuthConf.Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		return err
	}
	fmt.Println(token)
	//1. 토큰 검증하고

	//2. db 저장 유저 정보 업데이트

	//3. 토큰 생성

	return c.JSON(http.StatusOK, true)
}
