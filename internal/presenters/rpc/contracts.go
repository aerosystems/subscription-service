package RpcServer

import (
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
)

type SubscriptionUsecase interface {
	CreateFreeTrial(userUuid uuid.UUID, subscriptionType models.SubscriptionType) error
	GetSubscription(userUuid uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(userUuid uuid.UUID) error
}
