package subscription

import (
	"github.com/aerosystems/subscription-service/internal/presenters/http/handlers"
)

type Handler struct {
	*handlers.BaseHandler
	subscriptionUsecase handlers.SubscriptionUsecase
}

func NewSubscriptionHandler(baseHandler *handlers.BaseHandler, subscriptionUsecase handlers.SubscriptionUsecase) *Handler {
	return &Handler{baseHandler, subscriptionUsecase}
}
