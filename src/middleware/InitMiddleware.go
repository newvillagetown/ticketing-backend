package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/rs/zerolog"
	"main/common/logging"
	"os"
	"time"
)

func InitMiddleware(e *echo.Echo) error {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.URI == "/health" {
				return nil
			}
			startTime := time.Now()
			req := c.Request()
			url := req.URL.Path
			fmt.Println(url)
			if req.Method == "GET" && url == "/health" {
				return nil
			}
			rID := random.String(32)
			c.Set("rID", rID)

			//로그 데이터 생성한다.
			logData := logging.Log{}
			//TODO 추후 userID 어떤식으로 관리할지 정해지면 그때 넣는걸로
			logData.MakeLog("", url, req.Method, startTime, c.Response().Status, rID)
			//여기 나와야 로그를 찍을 수 있으니 로그 데이터를 만든다.
			//로그 찍는다. (local 일때만 찍는걸로 데이터 쌓이면 돈 많이 나가니..)
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")
			fmt.Println("??")
			return nil
		},
	}))
	return nil
}
