package handlers

import (
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/aerosystems/subs-service/pkg/monobank"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) WebhookMonobank(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*services.AccessTokenClaims)
	_ = accessTokenClaims
	paymentMethod := c.Param("payment_method")
	_ = paymentMethod
	xSignBase64 := c.Request().Header.Get("X-Sign")
	_ = xSignBase64
	var requestBody monobank.Webhook
	if err := c.Bind(&requestBody); err != nil {
		return h.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}
	return h.SuccessResponse(c, http.StatusOK, "webhook processed successfully", nil)
}
