package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"io/ioutil"
	"main/common/dbCommon/mongodbCommon"
	"main/common/errorCommon"
	"main/common/oauthCommon/google"
	"main/features/oauth/google/model"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
	_interface "main/features/oauth/google/usecase/interface"
	"main/middleware"
	"net/http"
	"time"
)

type CallbackGoogleOAuthHandler struct {
	UseCase _interface.ICallbackGoogleOAuthUseCase
}

func NewCallbackGoogleOAuthHandler() *CallbackGoogleOAuthHandler {
	return &CallbackGoogleOAuthHandler{UseCase: usecase.NewCallbackGoogleOAuthUseCase(repository.NewCallbackGoogleOAuthRepository(mongodbCommon.TokenCollection))}
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
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (cc *CallbackGoogleOAuthHandler) GoogleSignInCallback(c echo.Context) error {
	session, _ := middleware.Store.Get(c.Request(), "session")
	state := session.Values["state"]

	delete(session.Values, "state")
	session.Save(c.Request(), c.Response())
	if state != c.FormValue("state") {
		return errorCommon.ErrorMsg(errorCommon.ErrAuthFailed, errorCommon.Trace(), model.ErrAuthFailed, errorCommon.ErrFromInternal)
	}
	//인증서버에 액세스 토큰 요청
	token, err := google.OAuthConf.Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), err.Error(), errorCommon.ErrFromInternal)
	}

	client := google.OAuthConf.Client(oauth2.NoContext, token)
	userInfoResp, err := client.Get(google.UserInfoAPIEndpoint)
	if err != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), err.Error(), errorCommon.ErrFromInternal)
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), err.Error(), errorCommon.ErrFromInternal)
	}
	var authUser google.User
	json.Unmarshal(userInfo, &authUser)

	accessToken, _, err := cc.UseCase.CallbackGoogle(authUser)
	if err != nil {
		return err
	}
	//쿠키 셋팅
	cookie := &http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = accessToken
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, true)
}
