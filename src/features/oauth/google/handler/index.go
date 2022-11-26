package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
)

func IndexGoogleOAuthHandler(e *echo.Group) {
	handler := NewGoogleOAuthHandler()
	gApiV01Google := e.Group("/google")
	gApiV01Google.GET("/signin", handler.SignInGoogleOAuthHandler.SignInGoogle)
	gApiV01Google.GET("/signin/callback", handler.CallbackGoogleOAuthHandler.GoogleSignInCallback)

	gApiV01Google.Use(middleware.JWTWithConfig(jwtCommon.JwtConfig))
	gApiV01Google.GET("/signout", handler.SignOutGoogleOAuthHandler.SignOutGoogle)
}
