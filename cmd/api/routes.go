package main

import "github.com/labstack/echo/v4"

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/v1/subscriptions", app.BaseHandler.GetSubscriptions, app.AuthTokenMiddleware([]string{"user", "admin", "support"}))
	return e
}
