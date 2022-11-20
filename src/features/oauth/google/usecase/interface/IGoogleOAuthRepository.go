package _interface

import (
	"main/common/dbCommon/mongodb"
	"main/common/oauthCommon/google"
)

type ISignInGoogleOAuthRepository interface {
	SignInGoogle() error
}

type ICallbackGoogleOAuthRepository interface {
	CallbackGoogle() error
	CreateRefreshToken(token mongodb.RefreshToken) error
	DeleteAllRefreshToken(authUser google.User) error
}
