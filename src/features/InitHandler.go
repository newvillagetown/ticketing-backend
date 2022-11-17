package features

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	googleOAuthHandler "main/features/oauth/google/handler"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	gApiV01 := e.Group("/v0.1")

	gApiAuthV01 := gApiV01.Group("/auth")
	googleOAuthHandler.RegisterGoogleOAuthHandler(gApiAuthV01)

	gApiV01Features := gApiV01.Group("/features")
	gApiV01Features.Use(middleware.JWTWithConfig(jwtCommon.JwtConfig))
	return nil
}
