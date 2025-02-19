package GRPCServer

import (
	"context"
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/google/uuid"
)

type SubscriptionUsecase interface {
	CreateFreeTrial(ctx context.Context, customerUUID string) (*entities.Subscription, error)
	GetSubscription(ctx context.Context, customerUUID uuid.UUID) (*entities.Subscription, error)
	DeleteSubscription(ctx context.Context, customerUUID uuid.UUID) error
}
