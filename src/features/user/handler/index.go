package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/features/user/repository"
	"main/features/user/usecase"
)

func NewUserHandler(c *echo.Echo) {
	NewWithdrawalUserHandler(c, usecase.NewWithdrawalUserUseCase(repository.NewWithdrawalUserRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
}
