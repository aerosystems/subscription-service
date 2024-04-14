package usecases

import (
	"context"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
)

type InvoiceRepository interface {
	Create(invoice *models.Invoice) error
	GetByUserUuid(userUuid uuid.UUID) ([]models.Invoice, error)
	GetById(id int) (*models.Invoice, error)
	GetByAcquiringInvoiceId(acquiringInvoiceId string) (*models.Invoice, error)
	Update(invoice *models.Invoice) error
	Delete(invoice *models.Invoice) error
}

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *models.Subscription) error
	GetByUserUuid(ctx context.Context, userUuid uuid.UUID) (*models.Subscription, error)
	Update(ctx context.Context, subscription *models.Subscription) error
	Delete(ctx context.Context, subscription *models.Subscription) error
}

type PriceRepository interface {
	GetPrice(kindSubscription models.KindSubscription, durationSubscription models.DurationSubscription) (int, error)
	GetAll() map[models.KindSubscription]map[models.DurationSubscription]int
}
