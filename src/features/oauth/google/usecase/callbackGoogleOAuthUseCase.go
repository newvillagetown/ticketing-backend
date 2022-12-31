package usecase

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
	"main/common/jwtCommon"
	_interface "main/features/oauth/google/usecase/interface"
	"net/http"
	"time"
)

type CallbackGoogleOAuthUseCase struct {
	Repository     _interface.ICallbackGoogleOAuthRepository
	ContextTimeout time.Duration
}

func NewCallbackGoogleOAuthUseCase(repo _interface.ICallbackGoogleOAuthRepository, timeout time.Duration) _interface.ICallbackGoogleOAuthUseCase {
	return &CallbackGoogleOAuthUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

// TODO 트랜잭션 처리 필요
func (cc *CallbackGoogleOAuthUseCase) CallbackGoogle(c context.Context, user mysqlCommon.GormUser) (*http.Cookie, error) {
	ctx, cancel := context.WithTimeout(c, cc.ContextTimeout)
	defer cancel()

	now := time.Now()
	var userID string
	var userDTO mysqlCommon.GormUser
	//3. 유저 정보 저장(mysql db 저장)
	//db에 유저 정보가 있는지 체크
	userDTO, err := cc.Repository.FindOneUser(ctx, user)
	if err != nil {
		return nil, err
	}
	//없다면 유저 정보를 생성한다. (user, userAuth 테이블에 생성)
	if userDTO.GormModel.ID == "" {
		//TODO 트랜잭션 처리 필요
		userDTO = CreateMysqlUserDTO(user)
		userID = userDTO.GormModel.ID
		err = cc.Repository.CreateUser(ctx, userDTO)
		if err != nil {
			return nil, err
		}
		userAuthDTO := CreateMysqlUserAuthDTO(userDTO)
		err = cc.Repository.CreateUserAuth(ctx, userAuthDTO)
		if err != nil {
			return nil, err
		}
	}
	//1. 토큰 생성
	accessToken, refreshToken, err := jwtCommon.GenerateToken(userDTO.Email, now, userID)
	if err != nil {
		return nil, err
	}
	// 토큰 만들기
	token := CreateRefreshToken(userDTO, refreshToken, now)
	//기존 리프레시 토큰 제거
	err = cc.Repository.DeleteAllRefreshToken(ctx, userDTO)
	if err != nil {
		return nil, err
	}

	//2. 리프레시 토큰 저장
	err = cc.Repository.CreateRefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}
	cookie := CreateCookie(accessToken)

	return cookie, nil
}
