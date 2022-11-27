package handler

import (
	"github.com/labstack/echo/v4"
	_interface "main/features/product/usecase/interface"
)

type IRegisterProductHandler interface {
	Post(c echo.Context) error
}
type IGetProductHandler interface {
	Get(c echo.Context) error
}
type IGetsProductHandler interface {
	Gets(c echo.Context) error
}
type IDeleteProductHandler interface {
	Delete(c echo.Context) error
}
type IUpdateProductHandler interface {
	Update(c echo.Context) error
}
type IProductHandler interface {
	NewRegisterProductHandler(UseCase _interface.IRegisterProductUseCase) *RegisterProductHandler
	NewGetProductHandler(UseCase _interface.IGetProductUseCase) *GetProductHandler
	NewGetsProductHandler(UseCase _interface.IGetsProductUseCase) *GetsProductHandler
	NewDeleteProductHandler(UseCase _interface.IDeleteProductUseCase) *DeleteProductHandler
	NewUpdateProductHandler(UseCase _interface.IUpdateProductUseCase) *UpdateProductHandler
}

type ProductHandler struct {
	RegisterProductHandler RegisterProductHandler
	GetProductHandler      GetProductHandler
	GetsProductHandler     GetsProductHandler
	DeleteProductHandler   DeleteProductHandler
	UpdateProductHandler   UpdateProductHandler
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		RegisterProductHandler: *NewRegisterProductHandler(),
		GetProductHandler:      *NewGetProductHandler(),
		GetsProductHandler:     *NewGetsProductHandler(),
		DeleteProductHandler:   *NewDeleteProductHandler(),
		UpdateProductHandler:   *NewUpdateProductHandler(),
	}
}
