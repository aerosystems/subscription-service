package models

import (
	"github.com/google/uuid"
	"time"
)

type Invoice struct {
	Id            int           `json:"id"`
	Amount        int           `json:"amount"`
	UserUuid      uuid.UUID     `json:"userUuid"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	CreatedAt     time.Time     `json:"createdAtt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}

type InvoiceRepository interface {
	Create(invoice *Invoice) error
	GetByUserUuid(userUuid uuid.UUID) ([]Invoice, error)
	GetById(id int) (*Invoice, error)
	Update(invoice *Invoice) error
	Delete(invoice *Invoice) error
}
