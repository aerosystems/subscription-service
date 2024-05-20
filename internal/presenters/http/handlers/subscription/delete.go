package subscription

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (sh Handler) DeleteSubscription(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "not implemented")
}
