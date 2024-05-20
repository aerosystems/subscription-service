package payment

import (
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers"
)

type Handler struct {
	*handlers.BaseHandler
	paymentUsecase handlers.PaymentUsecase
}

func NewPaymentHandler(
	baseHandler *handlers.BaseHandler,
	paymentUsecase handlers.PaymentUsecase,
) *Handler {
	return &Handler{
		BaseHandler:    baseHandler,
		paymentUsecase: paymentUsecase,
	}
}
