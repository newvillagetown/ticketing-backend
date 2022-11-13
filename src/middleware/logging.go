package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"github.com/rs/zerolog/log"
	"time"
)

func RestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()
		req := c.Request()
		url := req.URL.Path
		c.Set("rID", random.String(32))
		err := next(c)

		if err != nil {
			//에러 로그 생성
			log.Error().Err(err).Msg("hi")
		} else {
			//일반 로그 생성
			log.Info().Str("URI", url).Int("state", 200).Int64("latency", time.Since(startTime).Milliseconds()).Str("METHOD", "POST").Msg("hi")
		}

		return nil
	}
}
