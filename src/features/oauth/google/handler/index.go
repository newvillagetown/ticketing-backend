package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
)

func NewGoogleOAuthHandler(c *echo.Echo) {
	NewSignInGoogleOAuthHandler(c, usecase.NewSignInGoogleOAuthUseCase(repository.NewSignInGoogleOAuthRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
	NewCallbackGoogleOAuthHandler(c, usecase.NewCallbackGoogleOAuthUseCase(repository.NewCallbackGoogleOAuthRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
	NewSignOutGoogleOAuthHandler(c, usecase.NewSignOutGoogleOAuthUseCase(repository.NewSignOutGoogleOAuthRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))

}
