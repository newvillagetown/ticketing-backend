package usecase

import (
	"main/common/jwtCommon"
	"main/common/oauthCommon/google"
	_interface "main/features/oauth/google/usecase/interface"
)

type CallbackGoogleOAuthUseCase struct {
	Repository _interface.ICallbackGoogleOAuthRepository
}

func NewCallbackGoogleOAuthUseCase(repo _interface.ICallbackGoogleOAuthRepository) _interface.ICallbackGoogleOAuthUseCase {
	return &CallbackGoogleOAuthUseCase{
		Repository: repo,
	}
}

func (cc *CallbackGoogleOAuthUseCase) CallbackGoogle(authUser google.User) (string, string, error) {
	//1. 토큰 생성
	accessToken, refreshToken, err := jwtCommon.GenerateToken(authUser.Email)
	if err != nil {
		return "", "", err
	}

	//2. 리프레시 토큰 저장

	//3. 유저 정보 저장

	return accessToken, refreshToken, nil
}
