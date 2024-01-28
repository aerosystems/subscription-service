package pg

import (
	"errors"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepo(db *gorm.DB) *InvoiceRepo {
	return &InvoiceRepo{db: db}
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
