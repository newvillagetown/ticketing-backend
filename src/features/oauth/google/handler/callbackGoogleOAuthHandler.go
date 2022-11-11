package handler

import (
	"github.com/labstack/echo/v4"
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
// @Success 200 {object} resGoogleSignin
// @Failure 400 {object} data.ErrResponse
// @Failure 500 {object} data.ErrResponse
// @Tags auth
func (cc *CallbackGoogleOAuthHandler) GoogleSignInCallback(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}
