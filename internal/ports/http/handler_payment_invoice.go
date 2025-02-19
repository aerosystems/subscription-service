package HTTPServer

import (
	"github.com/aerosystems/subscription-service/internal/entities"
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
func (h Handler) CreateInvoice(c echo.Context) error {
	user, err := GetUserFromContext(c.Request().Context())
	if err != nil {
		return err
	}

	method := entities.NewPaymentMethod(c.Param("payment_method"))

	var requestBody InvoiceRequest
	if err = c.Bind(&requestBody); err != nil {
		return entities.ErrInvalidRequestBody
	}

	paymentUrl, err := h.paymentUsecase.GetPaymentUrl(
		c.Request().Context(),
		user.UUID,
		method,
		entities.SubscriptionTypeFromString(requestBody.SubscriptionType),
		entities.SubscriptionDurationFromString(requestBody.SubscriptionDuration))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, InvoiceResponse{PaymentUrl: paymentUrl})
}
