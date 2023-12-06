package handlers

import (
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) CreateInvoice(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*services.AccessTokenClaims)
	_ = accessTokenClaims
	method := c.Param("payment_method")
	if err := h.paymentService.SetPaymentMethod(method); err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, "invalid payment method", err)
	}
	h.paymentService.CreateInvoice(nil)
	return c.JSON(http.StatusNotImplemented, "not implemented")
}
