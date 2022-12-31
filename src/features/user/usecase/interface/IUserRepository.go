package _interface

import "context"

type IWithdrawalUserRepository interface {
	WithdrawalUser(ctx context.Context, userID string) error
	WithdrawalUserAuth(ctx context.Context, userID string) error
}
