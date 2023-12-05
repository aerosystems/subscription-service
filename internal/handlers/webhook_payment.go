package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) WebhookPayment(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "not implemented")
}
