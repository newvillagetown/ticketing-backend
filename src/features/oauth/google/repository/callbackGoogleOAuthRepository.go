package repository

import _interface "main/features/oauth/google/usecase/interface"

func NewCallbackGoogleOAuthRepository() _interface.ICallbackGoogleOAuthRepository {
	return &CallbackGoogleOAuthRepository{}
}

func (cc *CallbackGoogleOAuthRepository) CallbackGoogle() error {

	return nil
}
