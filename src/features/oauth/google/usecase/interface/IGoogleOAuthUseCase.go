package _interface

import (
	"main/common/oauthCommon/google"
)

type ISignInGoogleOAuthUseCase interface {
	SignInGoogle() (string, string, error)
}
type ISignOutGoogleOAuthUseCase interface {
	SignOutGoogle(email string) error
}
type ICallbackGoogleOAuthUseCase interface {
	CallbackGoogle(authUser google.User) (string, string, error)
}
