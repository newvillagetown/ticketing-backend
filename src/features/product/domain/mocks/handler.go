package mocks

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	_ "main/features/product/usecase/interface"
	_interface "main/features/product/usecase/interface"
)

var (
	_ _interface.IGetsProductHandler   = &GetsProductHandler{}
	_ _interface.IGetProductHandler    = &GetProductHandler{}
	_ _interface.IDeleteProductHandler = &DeleteProductHandler{}
)

type DeleteProductHandler struct {
	mock.Mock
}
type GetsProductHandler struct {
	mock.Mock
}

type GetProductHandler struct {
	mock.Mock
}

func (d *DeleteProductHandler) Delete(c echo.Context) error {
	ret := d.Called(c)
	return ret.Error(1)
}

func (g *GetsProductHandler) Gets(c echo.Context) error {
	ret := g.Called(c)
	return ret.Error(1)
}

func (g *GetProductHandler) Get(c echo.Context) error {
	ret := g.Called(c, "dd")

	return ret.Error(1)
}
