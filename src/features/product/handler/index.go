package handler

import (
	"github.com/labstack/echo/v4"
)

func IndexProductHandler(e *echo.Group) {
	handler := NewProductHandler()
	gApiV01Features := e.Group("/product")

	gApiV01Features.POST("", handler.RegisterProductHandler.post)
	gApiV01Features.GET("", handler.GetProductHandler.Get)
	gApiV01Features.GET("/gets", handler.GetsProductHandler.gets)
	gApiV01Features.DELETE("", handler.DeleteProductHandler.delete)
	gApiV01Features.PUT("", handler.UpdateProductHandler.update)
}
