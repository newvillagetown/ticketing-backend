package features

import (
	"github.com/labstack/echo/v4"
	googleOAuthHandler "main/features/oauth/google/handler"
)

func InitHandler(e *echo.Echo) {
	gApiV01 := e.Group("/v0.1")
	googleOAuthHandler.RegisterGoogleOAuthHandler(gApiV01)
}
