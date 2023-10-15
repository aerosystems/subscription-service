package main

import (
	"errors"
	"github.com/aerosystems/subs-service/internal/handlers"
	AuthService "github.com/aerosystems/subs-service/pkg/auth_service"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

func (app *Config) AddMiddleware(e *echo.Echo, log *logrus.Logger) {
	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}
	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
	e.Use(middleware.Recover())
}

func (app *Config) AuthTokenMiddleware(roles []string) echo.MiddlewareFunc {
	AuthorizationConfig := echojwt.Config{
		SigningKey:     []byte(os.Getenv("ACCESS_SECRET")),
		ParseTokenFunc: parseToken,
		ErrorHandler: func(c echo.Context, err error) error {
			return handlers.ErrorResponse(c, http.StatusUnauthorized, "invalid token", err)
		},
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return AuthorizationConfig.ErrorHandler(c, errors.New("missing Authorization header"))
			}

			// Token should be in the form "Bearer <token>"
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return AuthorizationConfig.ErrorHandler(c, errors.New("invalid token format"))
			}

			token := tokenParts[1]

			accessTokenClaims, err := AuthService.DecodeAccessToken(token)
			if err != nil {
				return AuthorizationConfig.ErrorHandler(c, err)
			}

			if int64(accessTokenClaims.Exp) < time.Now().Unix() {
				return AuthorizationConfig.ErrorHandler(c, errors.New("token expired"))
			}

			// Перевірка, чи userRole знаходиться у переданому масиві roles
			roleFound := false
			for _, role := range roles {
				if accessTokenClaims.UserRole == role {
					roleFound = true
					break
				}
			}

			if !roleFound {
				// Якщо userRole не входить у roles, повертаємо помилку 403
				return handlers.ErrorResponse(c, http.StatusForbidden, "forbidden", errors.New("userRole not allowed"))
			}

			return next(c)
		}
	}
}

func (app *Config) BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, password, ok := c.Request().BasicAuth()

		if !ok || !checkCredentials(username, password) {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		return next(c)
	}
}

func checkCredentials(username, password string) bool {
	validUsername := os.Getenv("BASIC_AUTH_DOCS_USERNAME")
	validPassword := os.Getenv("BASIC_AUTH_DOCS_PASSWORD")

	return username == validUsername && password == validPassword
}

func parseToken(c echo.Context, auth string) (interface{}, error) {
	_ = c
	accessTokenClaims, err := AuthService.DecodeAccessToken(auth)
	if err != nil {
		return nil, err
	}

	return accessTokenClaims, nil
}
