package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
)

func IndexUserHandler(e *echo.Group) {
	handler := NewUserHandler()
	gApiV01User := e.Group("/user")
	gApiV01User.Use(middleware.JWTWithConfig(jwtCommon.JwtConfig))
	gApiV01User.POST("/withdrawal", handler.WithdrawalUserHandler.WithdrawalUser)
}
