package usecase

import _interface "main/features/oauth/google/usecase/interface"

type CallbackGoogleOAuthUseCase struct {
	Repository _interface.ICallbackGoogleOAuthRepository
}

func NewCallbackGoogleOAuthUseCase(repo _interface.ICallbackGoogleOAuthRepository) _interface.ICallbackGoogleOAuthUseCase {
	return &CallbackGoogleOAuthUseCase{
		Repository: repo,
	}
}

func (cc *CallbackGoogleOAuthUseCase) CallbackGoogle() error {

	return nil
}
