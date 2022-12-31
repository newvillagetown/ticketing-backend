package _interface

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
	"net/http"
)

type ISignInGoogleOAuthUseCase interface {
	SignInGoogle(c context.Context) error
}
type ISignOutGoogleOAuthUseCase interface {
	SignOutGoogle(c context.Context, email string) error
}
type ICallbackGoogleOAuthUseCase interface {
	CallbackGoogle(c context.Context, authUser mysqlCommon.GormUser) (*http.Cookie, error)
}
