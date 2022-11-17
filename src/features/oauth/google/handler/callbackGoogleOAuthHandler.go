package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"io/ioutil"
	"main/common/oauthCommon/google"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
	_interface "main/features/oauth/google/usecase/interface"
	"main/middleware"
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
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (cc *CallbackGoogleOAuthHandler) GoogleSignInCallback(c echo.Context) error {
	session, _ := middleware.Store.Get(c.Request(), "session")
	state := session.Values["state"]
	fmt.Println(state)

	delete(session.Values, "state")
	session.Save(c.Request(), c.Response())
	if state != c.FormValue("state") {
		return fmt.Errorf("Invalid session state")
	}
	//인증서버에 액세스 토큰 요청
	token, err := google.OAuthConf.Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		return err
	}

	//1. 토큰 검증하고
	client := google.OAuthConf.Client(oauth2.NoContext, token)
	userInfoResp, err := client.Get(google.UserInfoAPIEndpoint)
	if err != nil {
		return err
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		return err
	}
	var authUser google.User
	json.Unmarshal(userInfo, &authUser)
	//2. dbCommon 저장 유저 정보 업데이트
	fmt.Println(authUser.Name)
	fmt.Println(authUser.Email)

	//3. 토큰 생성

	return c.JSON(http.StatusOK, true)
}
