package services

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
)

type TokenService interface {
	GetAccessSecret() string
	DecodeAccessToken(tokenString string) (*AccessTokenClaims, error)
}

type SubsService interface {
	GetPrices() map[models.KindSubscription]map[models.DurationSubscription]int
	CreateFreeTrial(userUuid uuid.UUID, kind models.KindSubscription) error
	GetSubscription(userUuid uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(userUuid uuid.UUID) error
}

type PaymentService interface {
	SetPaymentMethod(paymentMethod string) error
	GetPaymentUrl(userUuid uuid.UUID, subscription models.KindSubscription, duration models.DurationSubscription) (string, error)
	ProcessingWebhookPayment(userUuid uuid.UUID, invoiceUuid uuid.UUID, acquiringInvoiceId string) error
}
