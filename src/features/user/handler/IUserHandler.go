package handler

import (
	"github.com/labstack/echo/v4"
	_interface "main/features/user/usecase/interface"
)

type IWithdrawalUserHandler interface {
	WithdrawalUser(c echo.Context) error
}
type IUserHandler interface {
	NewWithdrawalUserHandler(UseCase _interface.IWithdrawalUserUseCase) *WithdrawalUserHandler
}

type UserHandler struct {
	WithdrawalUserHandler WithdrawalUserHandler
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		WithdrawalUserHandler: *NewWithdrawalUserHandler(),
	}
}
