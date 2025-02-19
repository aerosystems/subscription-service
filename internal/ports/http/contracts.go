package HTTPServer

import (
	"context"
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/google/uuid"
)

type SubscriptionUsecase interface {
	CreateSubscription(ctx context.Context, userUuidStr, subscriptionTypeStr, subscriptionDurationStr string) (*entities.Subscription, error)
	CreateFreeTrial(ctx context.Context, userUuidStr string) (*entities.Subscription, error)
	GetSubscription(ctx context.Context, userUuid uuid.UUID) (*entities.Subscription, error)
	DeleteSubscription(ctx context.Context, userUuid uuid.UUID) error
}

type PaymentUsecase interface {
	GetPaymentUrl(ctx context.Context, userUuid uuid.UUID, paymentMethod entities.PaymentMethod, subscription entities.SubscriptionType, duration entities.SubscriptionDuration) (string, error)
	ProcessingWebhookPayment(ctx context.Context, paymentMethod entities.PaymentMethod, bodyBytes []byte, headers map[string][]string) error
	GetPrices(ctx context.Context) map[entities.SubscriptionType]map[entities.SubscriptionDuration]int
}
