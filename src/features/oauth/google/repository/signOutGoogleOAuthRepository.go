package repository

import _interface "main/features/oauth/google/usecase/interface"

func NewSignOutGoogleOAuthRepository() _interface.ISignOutGoogleOAuthRepository {
	return &SignOutGoogleOAuthRepository{}
}

func (s *SignOutGoogleOAuthRepository) SignOutGoogle() error {

	return nil
}
