package pg

import (
	"errors"
	"github.com/aerosystems/subs-service/internal/models"
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

type Invoice struct {
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

func (i Invoice) ToModel() models.Invoice {
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

func ModelToInvoiceFire(invoice models.Invoice) Invoice {
	return Invoice{
		Id:                 invoice.Id,
		Amount:             invoice.Amount,
		UserUuid:           invoice.UserUuid,
		InvoiceUuid:        invoice.InvoiceUuid,
		PaymentMethod:      invoice.PaymentMethod.String(),
		AcquiringInvoiceId: invoice.AcquiringInvoiceId,
		PaymentStatus:      invoice.PaymentStatus.String(),
		CreatedAt:          invoice.CreatedAt,
		UpdatedAt:          invoice.UpdatedAt,
	}
}

func (i *InvoiceRepo) Create(invoice *models.Invoice) error {
	result := i.db.Create(&invoice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (i *InvoiceRepo) GetByUserUuid(userUuid uuid.UUID) ([]models.Invoice, error) {
	var invoice []models.Invoice
	result := i.db.Where("user_uuid = ?", userUuid).Find(&invoice)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoice, nil
}

func (i *InvoiceRepo) GetById(id int) (*models.Invoice, error) {
	var invoice models.Invoice
	result := i.db.Where("id = ?", id).First(&invoice)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &invoice, nil
}

func (i *InvoiceRepo) GetByAcquiringInvoiceId(acquiringInvoiceId string) (*models.Invoice, error) {
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

func (i *InvoiceRepo) Update(invoice *models.Invoice) error {
	result := i.db.Save(&invoice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (i *InvoiceRepo) Delete(invoice *models.Invoice) error {
	result := i.db.Delete(&invoice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
