package usecase

import (
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

func (w *WithdrawalUserUseCase) WithdrawalUser() error {

	return nil
}
