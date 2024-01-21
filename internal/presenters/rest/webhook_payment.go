package rest

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (h *BaseHandler) WebhookPayment(c echo.Context) error {
	method := models.PaymentMethod(c.Param("payment_method"))
	if err := h.paymentService.SetPaymentMethod(method); err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, "invalid payment method", err)
	}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}
	if err := h.paymentService.ProcessingWebhookPayment(body, c.Request().Header); err != nil {
		return h.ErrorResponse(c, http.StatusInternalServerError, "could not process webhook", err)
	}
	return h.SuccessResponse(c, http.StatusOK, "webhook processed successfully", nil)
}
