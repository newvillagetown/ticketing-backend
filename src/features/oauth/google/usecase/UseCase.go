package usecase

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	"main/common/oauthCommon/google"
	"main/features/oauth/google/model/response"
	"net/http"
	"time"
)

func CallbackGoogleOAuthConvertRes(accessToken, refreshToken string) (response.ResCallbackGoogleOAuth, error) {
	if accessToken != "" || refreshToken != "" {
		return response.ResCallbackGoogleOAuth{}, fmt.Errorf("토큰이 비어있다.")
	}
	result := response.ResCallbackGoogleOAuth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return result, nil
}

func CreateRefreshToken(authUser mysqlCommon.GormUser, refreshToken string, now time.Time) mongodbCommon.RefreshToken {
	result := mongodbCommon.RefreshToken{
		Token:     refreshToken,
		Email:     authUser.Email,
		Created:   now,
		IsDeleted: false,
	}
	return result
}

func CreateMysqlUserDTO(authUser mysqlCommon.GormUser) mysqlCommon.GormUser {
	id := mysqlCommon.PKIDGenerate()
	authUser.GormModel.ID = id
	return authUser
}

func CreateMysqlUserAuthDTO(userDTO mysqlCommon.GormUser) mysqlCommon.GormUserAuth {
	id := mysqlCommon.PKIDGenerate()
	result := mysqlCommon.GormUserAuth{
		GormModel:  mysqlCommon.GormModel{ID: id},
		Provider:   "google",
		UserID:     userDTO.GormModel.ID,
		LastSignIn: time.Now().Unix(),
	}
	return result
}

func CreateCookie(accessToken string) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = accessToken
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(1 * time.Hour)
	return cookie
}

func CallGoogleOAuth(code string) (mysqlCommon.GormUser, error) {
	token, err := google.OAuthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return mysqlCommon.GormUser{}, errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), err.Error(), errorCommon.ErrFromInternal)
	}

	client := google.OAuthConf.Client(oauth2.NoContext, token)
	userInfoResp, err := client.Get(google.UserInfoAPIEndpoint)
	if err != nil {
		return mysqlCommon.GormUser{}, errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), err.Error(), errorCommon.ErrFromInternal)
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		return mysqlCommon.GormUser{}, errorCommon.ErrorMsg(errorCommon.ErrInternalServer, errorCommon.Trace(), err.Error(), errorCommon.ErrFromInternal)
	}
	var authUser google.User
	json.Unmarshal(userInfo, &authUser)

	result := mysqlCommon.GormUser{
		Email: authUser.Email,
		Name:  authUser.Name,
	}

	return result, nil
}
