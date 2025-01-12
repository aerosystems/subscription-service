package GRPCServer

import (
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
)

type SubscriptionUsecase interface {
	CreateFreeTrial(customerUUID string) (*models.Subscription, error)
	GetSubscription(customerUUID uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(customerUUID uuid.UUID) error
}
