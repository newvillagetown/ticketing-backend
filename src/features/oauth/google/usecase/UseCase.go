package usecase

import (
	"fmt"
	"main/features/oauth/google/model/response"
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
