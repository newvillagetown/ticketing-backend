package handler

import (
	"github.com/labstack/echo/v4"
	_interface "main/features/product/usecase/interface"
)

type IRegisterProductHandler interface {
	Post(c echo.Context) error
}

type IProductHandler interface {
	NewRegisterProductHandler(UseCase _interface.IRegisterProductUseCase) *RegisterProductHandler
}

type ProductHandler struct {
	RegisterProductHandler RegisterProductHandler
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		RegisterProductHandler: *NewRegisterProductHandler(),
	}
}
