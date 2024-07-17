package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServiceApiKeyAuth struct {
	ApiKey string
}

func NewServiceApiKeyAuth(apiKey string) (*ServiceApiKeyAuth, error) {
	if apiKey == "" {
		return nil, errors.New("api key is required")
	}
	return &ServiceApiKeyAuth{
		ApiKey: apiKey,
	}, nil
}

func (a ServiceApiKeyAuth) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("X-Api-Key") != a.ApiKey {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid api key")
			}
			return next(c)
		}
	}
}
