package RpcServer

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
)

type SubscriptionUsecase interface {
	CreateFreeTrial(userUuid uuid.UUID, kind models.KindSubscription) error
	GetSubscription(userUuid uuid.UUID) (*models.Subscription, error)
	DeleteSubscription(userUuid uuid.UUID) error
}
