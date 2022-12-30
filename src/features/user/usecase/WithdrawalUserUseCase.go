package usecase

import _interface "main/features/user/usecase/interface"

type WithdrawalUserUseCase struct {
	Repository _interface.IWithdrawalUserRepository
}

func NewWithdrawalUserUseCase(repo _interface.IWithdrawalUserRepository) _interface.IWithdrawalUserUseCase {
	return &WithdrawalUserUseCase{
		Repository: repo,
	}
}

// TODO 트랜잭션 처리 필요
func (w *WithdrawalUserUseCase) WithdrawalUser() error {
	return nil

}
