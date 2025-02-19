package HTTPServer

import (
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (h Handler) WebhookPayment(c echo.Context) error {
	method := entities.NewPaymentMethod(c.Param("payment_method"))
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return entities.ErrInvalidRequestPayload
	}
	if err := h.paymentUsecase.ProcessingWebhookPayment(c.Request().Context(), method, body, c.Request().Header); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
