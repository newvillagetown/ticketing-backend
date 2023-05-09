package _interface

import "context"

type IAuthTestUseCase interface {
	AuthTest(ctx context.Context) error
}
