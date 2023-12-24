package handlers

import (
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceRequest struct {
	Amount int `json:"amount" validate:"required,integer,gt=0" example:"1000"` // in cents
}

func (h *BaseHandler) CreateInvoice(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*services.AccessTokenClaims)
	method := c.Param("payment_method")
	if err := h.paymentService.SetPaymentMethod(method); err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, "invalid payment method", err)
	}
	h.paymentService.CreateInvoice(uuid.MustParse(accessTokenClaims.UserUuid), 5)
	return h.SuccessResponse(c, http.StatusCreated, "invoice successfully created", nil)
}
