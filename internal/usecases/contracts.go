package usecases

import (
	"context"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
)

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *models.Invoice) error
	GetByAcquiringInvoiceId(ctx context.Context, acquiringInvoiceId string) (*models.Invoice, error)
	Update(ctx context.Context, invoice *models.Invoice) error
}

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *models.Subscription) error
	GetByCustomerUuid(ctx context.Context, userUuid uuid.UUID) (*models.Subscription, error)
	Update(ctx context.Context, subscription *models.Subscription) error
	Delete(ctx context.Context, subscription *models.Subscription) error
}

type PriceRepository interface {
	GetPrice(kindSubscription models.SubscriptionType, durationSubscription models.SubscriptionDuration) (int, error)
	GetAll() map[models.SubscriptionType]map[models.SubscriptionDuration]int
}

type ProjectAdapter interface {
	PublishCreateProjectEvent(customerUuid string) error
}
