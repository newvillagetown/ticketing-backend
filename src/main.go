package main

import (
	_ "embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "main/docs"
	swaggerDocs "main/docs"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	swaggerDocs.SwaggerInfo.Host = "localhost:3000"
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.HideBanner = true
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "3000")))
}
