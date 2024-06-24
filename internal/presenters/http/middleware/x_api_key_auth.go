package middleware

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type XApiKeyAuth struct {
	ApiKey string
}

func NewXApiKeyAuth(apiKey string) (*XApiKeyAuth, error) {
	if apiKey == "" {
		return nil, errors.New("api key is required")
	}
	return &XApiKeyAuth{
		ApiKey: apiKey,
	}, nil
}

func (a XApiKeyAuth) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Printf("X-Api-Key: %s\n", c.Request().Header.Get("X-Api-Key"))
			if c.Request().Header.Get("X-Api-Key") != a.ApiKey {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid api key")
			}
			return next(c)
		}
	}
}
