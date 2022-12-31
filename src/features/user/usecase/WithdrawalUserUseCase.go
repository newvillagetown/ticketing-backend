package usecase

import (
	"context"
	_interface "main/features/user/usecase/interface"
	"time"
)

type WithdrawalUserUseCase struct {
	Repository     _interface.IWithdrawalUserRepository
	ContextTimeout time.Duration
}

func NewWithdrawalUserUseCase(repo _interface.IWithdrawalUserRepository, timeout time.Duration) _interface.IWithdrawalUserUseCase {
	return &WithdrawalUserUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (w *WithdrawalUserUseCase) WithdrawalUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, w.ContextTimeout)
	defer cancel()
	//TODO 트랜잭션 로직 적용 필요
	err := w.Repository.WithdrawalUser(ctx, userID)
	if err != nil {
		return err
	}
	err = w.Repository.WithdrawalUserAuth(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
