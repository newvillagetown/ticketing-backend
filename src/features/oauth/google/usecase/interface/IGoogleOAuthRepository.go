package _interface

import (
	"context"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
)

type ISignInGoogleOAuthRepository interface {
	SignInGoogle(ctx context.Context) error
}
type ISignOutGoogleOAuthRepository interface {
	SignOutGoogle(ctx context.Context) error
	DeleteRefreshToken(ctx context.Context, email string) error
}

type ICallbackGoogleOAuthRepository interface {
	CallbackGoogle(ctx context.Context) error
	CreateRefreshToken(ctx context.Context, token mongodbCommon.RefreshToken) error
	DeleteAllRefreshToken(ctx context.Context, userDTO mysqlCommon.GormUser) error
	FindOneUser(ctx context.Context, userDTO mysqlCommon.GormUser) (mysqlCommon.GormUser, error)
	CreateUser(ctx context.Context, userDTO mysqlCommon.GormUser) error
	CreateUserAuth(ctx context.Context, userAuthDTO mysqlCommon.GormUserAuth) error
}
