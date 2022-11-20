package usecase

import _interface "main/features/oauth/google/usecase/interface"

type SignOutGoogleOAuthUseCase struct {
	Repository _interface.ISignOutGoogleOAuthRepository
}

func NewSignOutGoogleOAuthUseCase(repo _interface.ISignOutGoogleOAuthRepository) _interface.ISignOutGoogleOAuthUseCase {
	return &SignOutGoogleOAuthUseCase{
		Repository: repo,
	}
}

func (s *SignOutGoogleOAuthUseCase) SignOutGoogle() error {

	return nil
}
