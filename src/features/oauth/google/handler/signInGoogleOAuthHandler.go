package handler

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/oauthCommon/google"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
	_interface "main/features/oauth/google/usecase/interface"
	"main/middleware"
	"net/http"
)

type SignInGoogleOAuthHandler struct {
	UseCase _interface.ISignInGoogleOAuthUseCase
}

func NewSignInGoogleOAuthHandler() *SignInGoogleOAuthHandler {
	return &SignInGoogleOAuthHandler{UseCase: usecase.NewSignInGoogleOAuthUseCase(repository.NewSignInGoogleOAuthRepository(mongodbCommon.TokenCollection))}
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
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (s *SignInGoogleOAuthHandler) SignInGoogle(c echo.Context) error {
	//TODO Google OAUTH 이슈
	sess, _ := middleware.Store.Get(c.Request(), "session")
	sess.Options = &sessions.Options{
		Path:     "/v0.1/auth/google/signin",
		MaxAge:   300,
		HttpOnly: true,
	}
	state := google.RandToken()
	sess.Values["state"] = state
	sess.Save(c.Request(), c.Response())
	c.Redirect(http.StatusMovedPermanently, google.GetLoginURL(state))
	return c.JSON(http.StatusOK, true)
}
