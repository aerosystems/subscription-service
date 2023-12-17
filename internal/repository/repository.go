package repository

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
)

type InvoiceRepository interface {
	Create(invoice *models.Invoice) error
	GetByUserUuid(userUuid uuid.UUID) ([]models.Invoice, error)
	GetById(id int) (*models.Invoice, error)
	Update(invoice *models.Invoice) error
	Delete(invoice *models.Invoice) error
}

type SubscriptionRepository interface {
	Create(subscription *models.Subscription) error
	GetByUserUuid(userUuid uuid.UUID) (*models.Subscription, error)
	Update(subscription *models.Subscription) error
	Delete(subscription *models.Subscription) error
}
