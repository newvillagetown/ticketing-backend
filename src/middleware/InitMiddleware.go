package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
)

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
	signingKey := jwtCommon.JwtSecretKey

	config := middleware.JWTConfig{
		TokenLookup: "query:token",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwtCommon signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			// claims are of type `jwtCommon.MapClaims` when token is created with `jwtCommon.Parse`
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

	e.Use(middleware.JWTWithConfig(config))

	return nil
}
