package main

import (
	_ "embed"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"main/common"
	"main/common/envCommon"
	_ "main/docs"
	swaggerDocs "main/docs"
	"main/features"
	mw "main/middleware"
)

func main() {
	if err := common.InitServer(); err != nil {
		fmt.Println(err)
		return
	}
	e := echo.New()

	//미들웨어 초기화
	err := mw.InitMiddleware(e)
	if err != nil {
		fmt.Println(err)
		return
	}
	//핸드러 초기화
	err = features.InitHandler(e)
	if err != nil {
		fmt.Println(err)
		return
	}

	// swagger 초기화
	if envCommon.Env.IsLocal {
		swaggerDocs.SwaggerInfo.Host = "localhost:3000"
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	} else {
		swaggerDocs.SwaggerInfo.Host = fmt.Sprintf("%s-%s.breathings.net", envCommon.Env.Env, envCommon.Env.Project)
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	e.HideBanner = true
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", envCommon.Env.Port)))
}
