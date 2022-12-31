package usecase

import (
	"context"
	_interface "main/features/oauth/google/usecase/interface"
	"time"
)

type SignOutGoogleOAuthUseCase struct {
	Repository     _interface.ISignOutGoogleOAuthRepository
	ContextTimeout time.Duration
}

func NewSignOutGoogleOAuthUseCase(repo _interface.ISignOutGoogleOAuthRepository, timeout time.Duration) _interface.ISignOutGoogleOAuthUseCase {
	return &SignOutGoogleOAuthUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (s *SignOutGoogleOAuthUseCase) SignOutGoogle(c context.Context, email string) error {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()
	err := s.Repository.DeleteRefreshToken(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
