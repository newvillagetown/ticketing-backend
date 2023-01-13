package features

import (
	"github.com/labstack/echo/v4"
	googleOAuthHandler "main/features/oauth/google/handler"
	productHandler "main/features/product/handler"
	userHandler "main/features/user/handler"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	//인증 핸들러 초기화
	googleOAuthHandler.NewGoogleOAuthHandler(e)

	//기능 핸들러 초기화
	productHandler.NewProductHandler(e)
	userHandler.NewUserHandler(e)

	return nil
}
