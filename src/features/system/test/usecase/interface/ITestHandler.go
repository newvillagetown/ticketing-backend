package _interface

import "github.com/labstack/echo/v4"

type IAuthTestHandler interface {
	AuthTest(c echo.Context) error
}
