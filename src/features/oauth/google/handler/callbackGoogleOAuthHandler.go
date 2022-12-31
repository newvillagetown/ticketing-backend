package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	"main/features/oauth/google/model"
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
	return &CallbackGoogleOAuthHandler{UseCase: usecase.NewCallbackGoogleOAuthUseCase(repository.NewCallbackGoogleOAuthRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut)}
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
	authUser, err := usecase.CallGoogleOAuth(c.FormValue("code"))
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	cookie, err := cc.UseCase.CallbackGoogle(ctx, authUser)
	if err != nil {
		return err
	}

	//쿠키 셋팅
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, true)
}
