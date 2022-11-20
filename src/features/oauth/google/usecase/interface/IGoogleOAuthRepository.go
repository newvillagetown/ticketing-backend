package _interface

import (
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/oauthCommon/google"
)

type ISignInGoogleOAuthRepository interface {
	SignInGoogle() error
}

type ICallbackGoogleOAuthRepository interface {
	CallbackGoogle() error
	CreateRefreshToken(token mongodbCommon.RefreshToken) error
	DeleteAllRefreshToken(authUser google.User) error
	FindOneUser(authUser google.User) (bool, error)
	CreateUser(userDTO mysqlCommon.User) error
	CreateUserAuth(userAuthDTO mysqlCommon.UserAuth) error
}
