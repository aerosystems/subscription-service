package payment

import (
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
// @Security XApiKeyAuth
// @Success 201 {object} Response{data=InvoiceResponse}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 422 {object} Response
// @Failure 500 {object} Response
// @Router /v1/invoices/{payment_method} [post]
func (ph Handler) CreateInvoice(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*middleware.AccessTokenClaims)
	method := models.NewPaymentMethod(c.Param("payment_method"))
	if err := ph.paymentUsecase.SetPaymentMethod(method); err != nil {
		return ph.ErrorResponse(c, http.StatusBadRequest, "invalid payment method", err)
	}
	var requestBody InvoiceRequest
	if err := c.Bind(&requestBody); err != nil {
		return ph.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}
	paymentUrl, err := ph.paymentUsecase.GetPaymentUrl(uuid.MustParse(accessTokenClaims.UserUuid), requestBody.SubscriptionType, requestBody.SubscriptionDuration)
	if err != nil {
		return ph.ErrorResponse(c, http.StatusInternalServerError, "could not create invoice", err)
	}
	return ph.SuccessResponse(c, http.StatusCreated, "invoice successfully created", InvoiceResponse{PaymentUrl: paymentUrl})
}
