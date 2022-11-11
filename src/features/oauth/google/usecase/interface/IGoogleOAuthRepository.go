package _interface

type ISignInGoogleOAuthRepository interface {
	SignInGoogle() error
}

type ICallbackGoogleOAuthRepository interface {
	CallbackGoogle() error
}
