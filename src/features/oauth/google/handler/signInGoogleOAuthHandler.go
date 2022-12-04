package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/features/oauth/google/model/response"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
	_interface "main/features/oauth/google/usecase/interface"
	"net/http"
	"time"
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
// @Success 200 {object} response.ResSignInGoogleOAuth
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (s *SignInGoogleOAuthHandler) SignInGoogle(c echo.Context) error {
	//TODO Google OAUTH 이슈
	/*
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
	*/

	accessToken, refreshToken, err := s.UseCase.SignInGoogle()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// cookie setting
	cookie := &http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = accessToken
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
	result := response.ResSignInGoogleOAuth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(http.StatusOK, result)
}
