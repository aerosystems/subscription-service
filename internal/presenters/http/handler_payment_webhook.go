package HTTPServer

import (
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (ph PaymentHandler) WebhookPayment(c echo.Context) error {
	method := models.NewPaymentMethod(c.Param("payment_method"))
	if err := ph.paymentUsecase.SetPaymentMethod(method); err != nil {
		return err
	}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return CustomErrors.ErrInvalidRequestPayload
	}
	if err := ph.paymentUsecase.ProcessingWebhookPayment(body, c.Request().Header); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
