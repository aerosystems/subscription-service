package payment

import (
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/aerosystems/subscription-service/internal/presenters/http/middleware"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceRequest struct {
	SubscriptionType     models.SubscriptionType     `json:"kindSubscription" validate:"required" example:"business"`
	SubscriptionDuration models.SubscriptionDuration `json:"durationSubscription" validate:"required" example:"12m"`
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
func (ph Handler) CreateInvoice(c echo.Context) error {
	user, err := middleware.GetUserFromContext(c.Request().Context())
	if err != nil {
		return err
	}
	method := models.NewPaymentMethod(c.Param("payment_method"))
	if err := ph.paymentUsecase.SetPaymentMethod(method); err != nil {
		return err
	}
	var requestBody InvoiceRequest
	if err := c.Bind(&requestBody); err != nil {
		return CustomErrors.ErrInvalidRequestBody
	}
	paymentUrl, err := ph.paymentUsecase.GetPaymentUrl(uuid.MustParse(user.Uuid), requestBody.SubscriptionType, requestBody.SubscriptionDuration)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, InvoiceResponse{PaymentUrl: paymentUrl})
}
