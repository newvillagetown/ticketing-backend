package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"main/common/errorCommon"
	"main/common/loggingCommon"
	"main/common/pubsubCommon"
	"strings"
	"time"
)

// RestLogger : log REST API
func RestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//로깅 초기 세팅
		startTime := time.Now()
		req := c.Request()
		url := req.URL.Path
		requestID := random.String(32)
		c.Set("requestID", requestID)
		if req.Method == "GET" && url == "/health" {
			return next(c)
		}

		err := next(c)
		//에러 파싱
		resError := errorCommon.Err{}
		var resCode int
		if c.Response().Status == 404 {
			err = errorCommon.ErrorMsg(errorCommon.ErrNotFound, "", fmt.Sprintf("Invalid url call : %s", url), errorCommon.ErrFromClient)
		}
		if err != nil {
			resError = ErrorParsing(err.Error())
			resCode = resError.HttpCode
		} else {
			resCode = c.Response().Status
		}

		// 로깅
		logging := loggingCommon.Log{}
		logging.MakeLog("", url, req.Method, startTime, resCode, requestID)
		if resCode >= 400 {
			//에러 로깅
			logging.MakeErrorLog(resError)
			loggingCommon.LogError(logging)
			//DB 부하를 생각해서 에러만 쌓는걸로
			pubsubCommon.PublishMessages(pubsubCommon.SubMongoDBLog, logging, pubsubCommon.PubSubCh)
			return echo.NewHTTPError(resError.HttpCode, errorCommon.ErrType(resError.ErrType).New(resError.ErrType, resError.Msg))
		} else {
			loggingCommon.LogInfo(logging)
		}
		return err
	}
}

func ErrorParsing(data string) errorCommon.Err {
	slice := strings.Split(data, "|")
	result := errorCommon.Err{
		HttpCode: errorCommon.ErrHttpCode[slice[0]],
		ErrType:  slice[0],
		Trace:    slice[1],
		Msg:      slice[2],
		From:     slice[3],
	}
	return result
}
