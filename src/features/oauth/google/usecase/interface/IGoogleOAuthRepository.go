package _interface

import "main/common/dbCommon/mongodb"

type ISignInGoogleOAuthRepository interface {
	SignInGoogle() error
}

type ICallbackGoogleOAuthRepository interface {
	CallbackGoogle() error
	CreateRefreshToken(token mongodb.RefreshToken) error
}
