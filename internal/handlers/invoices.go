package handlers

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceRequest struct {
	KindSubscription     models.KindSubscription     `json:"kindSubscription" validate:"required" example:"business"`
	DurationSubscription models.DurationSubscription `json:"durationSubscription" validate:"required" example:"annually"`
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
// @Param invoice body handlers.InvoiceRequest true "invoice"
// @Security ApiKeyAuth
// @Success 201 {object} Response{data=handlers.InvoiceResponse}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 422 {object} Response
// @Failure 500 {object} Response
// @Router /v1/invoices/{payment_method} [post]
func (h *BaseHandler) CreateInvoice(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*services.AccessTokenClaims)
	method := c.Param("payment_method")
	if err := h.paymentService.SetPaymentMethod(method); err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, "invalid payment method", err)
	}
	var requestBody InvoiceRequest
	if err := c.Bind(&requestBody); err != nil {
		return h.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}
	paymentUrl, err := h.paymentService.GetPaymentUrl(uuid.MustParse(accessTokenClaims.UserUuid), requestBody.KindSubscription, requestBody.DurationSubscription)
	if err != nil {
		return h.ErrorResponse(c, http.StatusInternalServerError, "could not create invoice", err)
	}
	return h.SuccessResponse(c, http.StatusCreated, "invoice successfully created", InvoiceResponse{PaymentUrl: paymentUrl})
}
