package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"io"
	"main/common/errorSystem"
	"main/common/logging"
	"time"
)

func RestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()
		requestParam := GetJSONRawBody(c)
		req := c.Request()
		url := req.URL.Path
		if req.Method == "GET" && url == "/health" {
			return next(c)
		}
		rID := random.String(32)
		c.Set("rID", rID)
		err := next(c)

		//로그 데이터 생성한다.
		logData := logging.Log{}
		//TODO 추후 userID 어떤식으로 관리할지 정해지면 그때 넣는걸로
		logData.MakeLog("", url, req.Method, startTime, c.Response().Status, rID)
		//여기 나와야 로그를 찍을 수 있으니 로그 데이터를 만든다.
		if err != nil {
			resError := errorSystem.ErrorParsing(err.Error())
			logData.MakeErrorLog(requestParam, resError)
			logData.ErrorLog()
		} else {
			//로그 찍는다. (local 일때만 찍는걸로 데이터 쌓이면 돈 많이 나가니..)
			logData.InfoLog()
		}

		//TODO 몽고디비 데이터 전송

		return nil
	}
}

func GetJSONRawBody(c echo.Context) map[string]interface{} {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Use the content
	jsonData := make(map[string]interface{})
	json.Unmarshal(bodyBytes, &jsonData)
	return jsonData
}
