package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ApiKeyAuth struct {
	ApiKey string
}

func NewApiKeyAuth(apiKey string) (*ApiKeyAuth, error) {
	if apiKey == "" {
		return nil, errors.New("api key is required")
	}
	return &ApiKeyAuth{
		ApiKey: apiKey,
	}, nil
}

func (a ApiKeyAuth) XApiKey() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("X-Api-Key") != a.ApiKey {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid api key")
			}
			return next(c)
		}
	}
}
