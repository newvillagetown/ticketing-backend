package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	swaggerDocs "main/docs"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	swaggerDocs.SwaggerInfo.Host = "localhost:3000"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "3000")))
}
