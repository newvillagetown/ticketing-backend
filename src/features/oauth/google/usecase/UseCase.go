package usecase

import (
	"fmt"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/oauthCommon/google"
	"main/features/oauth/google/model/response"
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

func CreateRefreshToken(authUser google.User, refreshToken string, now time.Time) mongodbCommon.RefreshToken {
	result := mongodbCommon.RefreshToken{
		Token:     refreshToken,
		Email:     authUser.Email,
		Created:   now,
		IsDeleted: false,
	}
	return result
}

func CreateMysqlUserDTO(authUser google.User) mysqlCommon.User {
	id := mysqlCommon.PKIDGenerate()
	now := time.Now().Format("2006-01-02 15:04:05")
	result := mysqlCommon.User{
		ID:        id,
		Email:     authUser.Email,
		Name:      authUser.Name,
		Created:   now,
		IsDeleted: false,
	}
	return result
}

func CreateMysqlUserAuthDTO(userDTO mysqlCommon.User) mysqlCommon.UserAuth {
	id := mysqlCommon.PKIDGenerate()
	now := time.Now().Format("2006-01-02 15:04:05")
	result := mysqlCommon.UserAuth{
		ID:         id,
		Provider:   "google",
		UserID:     userDTO.ID,
		LastSignIn: now,
		Created:    userDTO.Created,
		IsDeleted:  false,
	}
	return result
}
