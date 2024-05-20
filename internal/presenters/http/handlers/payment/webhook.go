package payment

import (
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (ph Handler) WebhookPayment(c echo.Context) error {
	method := models.NewPaymentMethod(c.Param("payment_method"))
	if err := ph.paymentUsecase.SetPaymentMethod(method); err != nil {
		return ph.ErrorResponse(c, http.StatusBadRequest, "invalid payment method", err)
	}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return ph.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}
	if err := ph.paymentUsecase.ProcessingWebhookPayment(body, c.Request().Header); err != nil {
		return ph.ErrorResponse(c, http.StatusInternalServerError, "could not process webhook", err)
	}
	return ph.SuccessResponse(c, http.StatusOK, "webhook processed successfully", nil)
}
