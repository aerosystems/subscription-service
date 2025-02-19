package usecases

import (
	"context"
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/google/uuid"
)

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *entities.Invoice) error
	GetByAcquiringInvoiceId(ctx context.Context, acquiringInvoiceId string) (*entities.Invoice, error)
	Update(ctx context.Context, invoice *entities.Invoice) error
}

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *entities.Subscription) error
	GetByCustomerUuid(ctx context.Context, userUuid uuid.UUID) (*entities.Subscription, error)
	Update(ctx context.Context, subscription *entities.Subscription) error
	Delete(ctx context.Context, subscriptionUUID uuid.UUID) error
}

type PriceRepository interface {
	GetPrice(subscriptionType entities.SubscriptionType, subscriptionDuration entities.SubscriptionDuration) (int, error)
	GetAll() map[entities.SubscriptionType]map[entities.SubscriptionDuration]int
}
