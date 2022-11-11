package usecase

import _interface "main/features/oauth/google/usecase/interface"

type SignInGoogleOAuthUseCase struct {
	Repository _interface.ISignInGoogleOAuthRepository
}

func NewSignInGoogleOAuthUseCase(repo _interface.ISignInGoogleOAuthRepository) _interface.ISignInGoogleOAuthUseCase {
	return &SignInGoogleOAuthUseCase{
		Repository: repo,
	}
}

func (s *SignInGoogleOAuthUseCase) SignInGoogle() error {

	return nil
}
