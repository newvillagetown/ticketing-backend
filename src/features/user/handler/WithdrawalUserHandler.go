package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/features/user/repository"
	"main/features/user/usecase"

	_interface "main/features/user/usecase/interface"
	"net/http"
)

type WithdrawalUserHandler struct {
	UseCase _interface.IWithdrawalUserUseCase
}

func NewWithdrawalUserHandler() *WithdrawalUserHandler {
	return &WithdrawalUserHandler{UseCase: usecase.NewWithdrawalUserUseCase(repository.NewWithdrawalUserRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut)}
}

// withdrawal user
// @Router /v0.1/auth/user/withdrawal [post]
// @Summary 회원 탈퇴
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (w *WithdrawalUserHandler) WithdrawalUser(c echo.Context) error {

	return c.JSON(http.StatusOK, true)
}
