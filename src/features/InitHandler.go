package features

import (
	"fmt"
	"github.com/labstack/echo/v4"
	googleOAuthHandler "main/features/oauth/google/handler"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/health/ok", func(c echo.Context) error {
		fmt.Println("여기 들어오나??")
		return c.JSON(http.StatusOK, "ok")
	})
	gApiV01 := e.Group("/google")
	googleOAuthHandler.RegisterGoogleOAuthHandler(gApiV01)
	return nil
}
