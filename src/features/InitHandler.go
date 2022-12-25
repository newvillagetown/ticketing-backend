package features

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	googleOAuthHandler "main/features/oauth/google/handler"
	productHandler "main/features/product/handler"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	gApiV01 := e.Group("/v0.1")

	//인증 핸들러 초기화
	gApiAuthV01 := gApiV01.Group("/auth")
	googleOAuthHandler.IndexGoogleOAuthHandler(gApiAuthV01)

	//기능별 핸들러 초기화
	gApiV01Features := gApiV01.Group("/features")
	gApiV01Features.Use(middleware.JWTWithConfig(jwtCommon.JwtConfig))

	productHandler.IndexProductHandler(gApiV01Features)

	//테스트
	gApiAuthV01.GET("/test", func(c echo.Context) error {
		fmt.Println(c.Cookie("hello"))
		return c.JSON(http.StatusOK, true)
	})
	return nil
}
