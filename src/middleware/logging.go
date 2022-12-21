package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/rs/zerolog"
	"main/common/errorCommon"
	"main/common/loggingCommon"
	"os"
	"time"
)

func RestLogger(c echo.Context, v middleware.RequestLoggerValues) error {
	logger := zerolog.New(os.Stdout)
	if v.URI == "/health" {
		return nil
	}
	startTime := time.Now()
	req := c.Request()
	url := req.URL.Path
	if req.Method == "GET" && url == "/health" {
		return nil
	}
	rID := random.String(32)
	c.Set("rID", rID)
	//로그 데이터 생성한다.
	logData := loggingCommon.Log{}
	if c.Response().Status >= 400 {
		fmt.Println("여기 들어올텐데")
		//에러 로그 처리
		logger.Info().Err(v.Error).
			Str("URI", v.URI).
			Int("status", errorCommon.GetStatusCode(v.Error)).
			Msg("request")
	} else {
		//엑세스 로그 처리
		logData.MakeLog("", url, req.Method, startTime, c.Response().Status, rID)
		logData.InfoLog()
		fmt.Println(logData)
	}

	return nil
}
