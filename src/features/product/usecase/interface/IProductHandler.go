package _interface

import (
	"github.com/labstack/echo/v4"
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
