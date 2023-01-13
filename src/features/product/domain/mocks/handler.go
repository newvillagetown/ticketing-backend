package mocks

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"main/features/product/handler"
)

var (
	_ handler.IGetsProductHandler = &GetsProductHandler{}
)

type GetsProductHandler struct {
	mock.Mock
}

func (r *GetsProductHandler) Gets(c echo.Context) error {
	ret := r.Called(c)
	return ret.Error(0)
}
