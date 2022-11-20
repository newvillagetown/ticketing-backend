package _interface

import (
	"main/common/dbCommon/mongodbCommon"
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
	CreateUser(authUser google.User) error
}
