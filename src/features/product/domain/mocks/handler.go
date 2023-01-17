package mocks

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"main/features/product/domain/response"
	_ "main/features/product/usecase/interface"
	_interface "main/features/product/usecase/interface"
)

var (
	_ _interface.IGetsProductHandler = &GetsProductHandler{}
)

type GetsProductHandler struct {
	mock.Mock
}

func (r *GetsProductHandler) Gets(c echo.Context) ([]*response.ResGetsProduct, error) {
	ret := r.Called(c)
	return ret.Get(0).([]*response.ResGetsProduct), ret.Error(1)
}
