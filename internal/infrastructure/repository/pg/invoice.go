package pg

import (
	"context"
	"errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type InvoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepo(db *gorm.DB) *InvoiceRepo {
	return &InvoiceRepo{db: db}
}

type InvoicePg struct {
	Id                 int       `gorm:"primaryKey;autoIncrement"`
	Amount             int       `gorm:"<-"`
	UserUuid           uuid.UUID `gorm:"<-"`
	InvoiceUuid        uuid.UUID `gorm:"unique"`
	PaymentMethod      string    `gorm:"<-"`
	AcquiringInvoiceId string    `gorm:"<-"`
	PaymentStatus      string    `gorm:"<-"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

func (i InvoicePg) ToModel() models.Invoice {
	return models.Invoice{
		Id:                 i.Id,
		Amount:             i.Amount,
		UserUuid:           i.UserUuid,
		InvoiceUuid:        i.InvoiceUuid,
		PaymentMethod:      models.NewPaymentMethod(i.PaymentMethod),
		AcquiringInvoiceId: i.AcquiringInvoiceId,
		PaymentStatus:      models.NewPaymentStatus(i.PaymentStatus),
		CreatedAt:          i.CreatedAt,
		UpdatedAt:          i.UpdatedAt,
	}
}

func (i *InvoiceRepo) Create(_ context.Context, invoice *models.Invoice) error {
	result := i.db.Create(&invoice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (i *InvoiceRepo) GetByAcquiringInvoiceId(_ context.Context, acquiringInvoiceId string) (*models.Invoice, error) {
	var invoice models.Invoice
	result := i.db.Where("acquiring_invoice_id = ?", acquiringInvoiceId).First(&invoice)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &invoice, nil
}

func (i *InvoiceRepo) Update(_ context.Context, invoice *models.Invoice) error {
	result := i.db.Save(&invoice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
