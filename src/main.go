package main

import (
	_ "embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"main/common"
	_ "main/docs"
	swaggerDocs "main/docs"
	"main/features"
	"net/http"
)

func main() {
	if err := common.InitEnv(); err != nil {
		fmt.Println("서버 에러 발생")
		return
	}
	if err := common.InitAws(); err != nil {
		fmt.Println("aws 초기화 에러")
		return
	}
	if err := common.GoogleOauthInit(); err != nil {
		fmt.Println("구글 초기화 에러")
		return
	}
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	//핸드러 초기화
	features.InitHandler(e)

	// swagger 초기화
	if common.Env.IsLocal {
		swaggerDocs.SwaggerInfo.Host = "localhost:3000"
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	} else {
		swaggerDocs.SwaggerInfo.Host = fmt.Sprintf("%s-%s.breathings.net", common.Env.Env, "ticketing")
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.HideBanner = true
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", common.Env.Port)))
}
