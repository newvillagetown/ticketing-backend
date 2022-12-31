package usecase

import (
	"context"
	_interface "main/features/oauth/google/usecase/interface"
	"time"
)

type SignInGoogleOAuthUseCase struct {
	Repository     _interface.ISignInGoogleOAuthRepository
	ContextTimeout time.Duration
}

func NewSignInGoogleOAuthUseCase(repo _interface.ISignInGoogleOAuthRepository, timeout time.Duration) _interface.ISignInGoogleOAuthUseCase {
	return &SignInGoogleOAuthUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (s *SignInGoogleOAuthUseCase) SignInGoogle(c context.Context) error {
	return nil
}
