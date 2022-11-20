package _interface

import (
	"main/common/oauthCommon/google"
)

type ISignInGoogleOAuthUseCase interface {
	SignInGoogle() error
}
type ISignOutGoogleOAuthUseCase interface {
	SignOutGoogle() error
}
type ICallbackGoogleOAuthUseCase interface {
	CallbackGoogle(authUser google.User) (string, string, error)
}
