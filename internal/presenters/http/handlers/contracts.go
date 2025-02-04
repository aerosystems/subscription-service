package handlers

import (
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
)

type SubscriptionUsecase interface {
	CreateSubscription(userUuidStr, subscriptionTypeStr, subscriptionDurationStr string) (*models.Subscription, error)
	CreateFreeTrial(userUuidStr string) (*models.Subscription, error)
	GetSubscription(userUuid uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(userUuid uuid.UUID) error
}

type PaymentUsecase interface {
	SetPaymentMethod(paymentMethod models.PaymentMethod) error
	GetPaymentUrl(userUuid uuid.UUID, subscription models.SubscriptionType, duration models.SubscriptionDuration) (string, error)
	ProcessingWebhookPayment(bodyBytes []byte, headers map[string][]string) error
	GetPrices() map[models.SubscriptionType]map[models.SubscriptionDuration]int
}
