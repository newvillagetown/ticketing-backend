package features

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	googleOAuthHandler "main/features/oauth/google/handler"
	productHandler "main/features/product/handler"
	"net/http"
	"time"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	gApiV01 := e.Group("/v0.1")

	gApiAuthV01 := gApiV01.Group("/auth")
	googleOAuthHandler.IndexGoogleOAuthHandler(gApiAuthV01)
	productHandler.IndexProductHandler(gApiV01)

	gApiV01Features := gApiV01.Group("/features")
	gApiV01Features.Use(middleware.JWTWithConfig(jwtCommon.JwtConfig))
	gApiV01.GET("/test", func(c echo.Context) error {
		fmt.Println("api 서버 타이암웃 24초 설정 테스트 시작 ")
		for i := 1; i <= 30; i++ {
			time.Sleep(1 * time.Second) // 1s
			fmt.Println(fmt.Sprintf("%d 초 ", i))
		}
		return c.JSON(http.StatusOK, true)
	})
	return nil
}
