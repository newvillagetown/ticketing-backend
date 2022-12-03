package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/dbCommon/mysqlCommon"
	"main/common/jwtCommon"
	"time"
)

var Store = sessions.NewCookieStore([]byte("secret"))

func InitMiddleware(e *echo.Echo) error {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:        true,
		LogStatus:     true,
		LogValuesFunc: RestLogger,
	}))
	signingKey := jwtCommon.AccessToknenSecretKey

	jwtCommon.JwtConfig = middleware.JWTConfig{
		TokenLookup: "header:token",
		Claims:      &jwtCommon.JwtCustomClaims{},
		SigningKey:  signingKey,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, fmt.Errorf("invalid token")
			}
			return token, nil
		},
	}
	e.Use(SetDBMiddleware)

	return nil
}

/*
Continuous session mode which might be helpful when handling API requests,
for example, you can set up *gorm.DB with Timeout Context in middlewares,
and then use the *gorm.DB when processing all requests
*/
func SetDBMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		timeoutContext, _ := context.WithTimeout(context.Background(), 8*time.Second)
		ctx := context.WithValue(c.Request().Context(), "DB", mysqlCommon.GormDB.WithContext(timeoutContext))
		c.Request().WithContext(ctx)
		return next(c)
	}
}
