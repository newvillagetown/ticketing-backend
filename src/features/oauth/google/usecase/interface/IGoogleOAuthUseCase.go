package _interface

type ISignInGoogleOAuthUseCase interface {
	SignInGoogle() error
}

type ICallbackGoogleOAuthUseCase interface {
	CallbackGoogle() error
}
