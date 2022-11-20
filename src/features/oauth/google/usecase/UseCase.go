package usecase

import (
	"fmt"
	"main/common/dbCommon/mongodbCommon"
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
