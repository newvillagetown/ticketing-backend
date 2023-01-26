package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/features/message/repository"
	"main/features/message/usecase"
)

func NewNaverSmsHandler(c *echo.Echo) {
	NewSendNaverSmsHandler(c, usecase.NewSendNaverSmsUseCase(repository.NewSendNaverSmsRepository(mysqlCommon.GormDB, mongodbCommon.EventCollection), mysqlCommon.DBTimeOut))
}
