package handlers

import (
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) WebhookPayment(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*services.AccessTokenClaims)
	_ = accessTokenClaims
	return c.JSON(http.StatusNotImplemented, "not implemented")
}
