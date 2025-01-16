package HTTPServer

import (
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/go-logrusutil/logrusutil/logctx"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceRequest struct {
	SubscriptionType     string `json:"subscriptionType" validate:"required" example:"business"`
	SubscriptionDuration string `json:"subscriptionDuration" validate:"required" example:"12m"`
}

type InvoiceResponse struct {
	PaymentUrl string `json:"paymentUrl" example:"https://api.monobank.ua"`
}

// CreateInvoice godoc
// @Summary Create invoice
// @Description Create invoice by payment method
// @Tags invoices
// @Accept  json
// @Produce  json
// @Param payment_method path string true "payment method" Enums(monobank)
// @Param invoice body InvoiceRequest true "invoice"
// @Security ServiceApiKeyAuth
// @Success 201 {object} InvoiceResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 401 {object} handlers.ErrorResponse
// @Failure 422 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /v1/invoices/{payment_method} [post]
func (ph PaymentHandler) CreateInvoice(c echo.Context) error {
	user, err := GetUserFromContext(c.Request().Context())
	if err != nil {
		return err
	}

	method := models.NewPaymentMethod(c.Param("payment_method"))
	logctx.From(c.Request().Context()).WithField("payment_method", method).Info("set payment method")
	if err = ph.paymentUsecase.SetPaymentMethod(method); err != nil {
		return err
	}

	var requestBody InvoiceRequest
	if err = c.Bind(&requestBody); err != nil {
		return CustomErrors.ErrInvalidRequestBody
	}

	paymentUrl, err := ph.paymentUsecase.GetPaymentUrl(
		user.UUID,
		models.SubscriptionTypeFromString(requestBody.SubscriptionType),
		models.SubscriptionDurationFromString(requestBody.SubscriptionDuration))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, InvoiceResponse{PaymentUrl: paymentUrl})
}
