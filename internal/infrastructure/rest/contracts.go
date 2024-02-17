package rest

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
)

type SubscriptionUsecase interface {
	GetPrices() map[models.KindSubscription]map[models.DurationSubscription]int
	CreateFreeTrial(userUuid uuid.UUID, kind models.KindSubscription) error
	GetSubscription(userUuid uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(userUuid uuid.UUID) error
}

type PaymentUsecase interface {
	SetPaymentMethod(paymentMethod models.PaymentMethod) error
	GetPaymentUrl(userUuid uuid.UUID, subscription models.KindSubscription, duration models.DurationSubscription) (string, error)
	ProcessingWebhookPayment(bodyBytes []byte, headers map[string][]string) error
}
