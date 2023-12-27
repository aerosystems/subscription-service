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
	DurationSubscription models.DurationSubscription `json:"durationSubscription" validate:"required" example:"12m"`
}

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
	h.paymentService.CreateInvoice(uuid.MustParse(accessTokenClaims.UserUuid), requestBody.KindSubscription, requestBody.DurationSubscription)
	return h.SuccessResponse(c, http.StatusCreated, "invoice successfully created", nil)
}
