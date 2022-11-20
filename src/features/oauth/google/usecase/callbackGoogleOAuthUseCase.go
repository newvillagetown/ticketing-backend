package usecase

import (
	"fmt"
	"main/common/jwtCommon"
	"main/common/oauthCommon/google"
	_interface "main/features/oauth/google/usecase/interface"
	"time"
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
	now := time.Now()
	//1. 토큰 생성
	accessToken, refreshToken, err := jwtCommon.GenerateToken(authUser.Email, now)
	if err != nil {
		return "", "", err
	}
	// 토큰 만들기
	token := CreateRefreshToken(authUser, refreshToken, now)
	fmt.Println(token)
	//기존 리프레시 토큰 제거
	err = cc.Repository.DeleteAllRefreshToken(authUser)
	if err != nil {
		return "", "", err
	}

	//2. 리프레시 토큰 저장
	err = cc.Repository.CreateRefreshToken(token)
	if err != nil {
		return "", "", err
	}

	//3. 유저 정보 저장(mysql db 저장)

	//db에 유저 정보가 있는지 체크
	//없다면 유저 정보를 생성한다.

	return accessToken, refreshToken, nil
}
