package features

import (
	"github.com/labstack/echo/v4"
	googleOAuthHandler "main/features/oauth/google/handler"
	"main/middleware"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/health/ok", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	gApiV01 := e.Group("/v0.1")
	gApiV01.Use(middleware.RestLogger)
	googleOAuthHandler.RegisterGoogleOAuthHandler(gApiV01)
	return nil
}
