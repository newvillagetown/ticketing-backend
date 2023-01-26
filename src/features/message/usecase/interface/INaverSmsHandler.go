package _interface

import "github.com/labstack/echo/v4"

type ISendNaverSmsHandler interface {
	Send(c echo.Context) error
}
