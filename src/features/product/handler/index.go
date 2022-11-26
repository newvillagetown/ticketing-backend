package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
)

func IndexProductHandler(e *echo.Group) {
	handler := NewProductHandler()
	gApiV01Features := e.Group("/features")

	gApiV01Features.Use(middleware.JWTWithConfig(jwtCommon.JwtConfig))
	gApiV01Features.POST("/product", handler.RegisterProductHandler.post)
	gApiV01Features.GET("/product", handler.GetProductHandler.get)
	gApiV01Features.GET("/product/gets", handler.GetsProductHandler.gets)
}
