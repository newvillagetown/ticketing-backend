package _interface

import "context"

type IWithdrawalUserUseCase interface {
	WithdrawalUser(c context.Context, userID string) error
}
