package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"main/common/logging"
	"time"
)

func RestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()
		req := c.Request()
		url := req.URL.Path
		rID := random.String(32)
		c.Set("rID", rID)
		err := next(c)

		//로그 데이터 생성한다.
		logData := logging.Log{}
		logData.MakeLog("", url, "", startTime, 200, rID)
		//여기 나와야 로그를 찍을 수 있으니 로그 데이터를 만든다.
		if err != nil {
			logData.MakeErrorLog()
			logData.ErrorLog()
		} else {
			//로그 찍는다. (local 일때만 찍는걸로 데이터 쌓이면 돈 많이 나가니..)
			logData.InfoLog()
		}

		return nil
	}
}
