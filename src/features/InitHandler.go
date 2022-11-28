package features

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	"main/common/s3Common"
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

	gApiAuthV01 := gApiV01.Group("/auth")
	googleOAuthHandler.IndexGoogleOAuthHandler(gApiAuthV01)
	productHandler.IndexProductHandler(gApiV01)

	gApiV01Features := gApiV01.Group("/features")
	gApiV01Features.Use(middleware.JWTWithConfig(jwtCommon.JwtConfig))
	gApiV01Features.GET("/test", func(c echo.Context) error {
		signedURL, err := s3Common.ImageGetSignedURL("infra.png", s3Common.ImgTypeProduct)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(signedURL)
		return c.JSON(http.StatusOK, true)
	})
	return nil
}
