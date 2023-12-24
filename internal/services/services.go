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
	CreateFreeTrial(userUuid uuid.UUID, kind models.KindSubscription) error
	GetSubscription(userUuid uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(userUuid uuid.UUID) error
}

type PaymentService interface {
	SetPaymentMethod(paymentMethod string) error
	CreateInvoice(userUuid uuid.UUID, amount int) (*models.Invoice, error)
}
