package repository

import (
	_interface "main/features/oauth/google/usecase/interface"
)

func NewSignInGoogleOAuthRepository() _interface.ISignInGoogleOAuthRepository {
	return &SignInGoogleOAuthRepository{}
}

func (s *SignInGoogleOAuthRepository) SignInGoogle() error {

	return nil
}
