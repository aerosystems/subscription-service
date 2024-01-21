package models

import (
	"github.com/google/uuid"
	"time"
)

type Invoice struct {
	Id                 int           `json:"id" gorm:"primaryKey;autoIncrement"`
	Amount             int           `json:"amount"`
	UserUuid           uuid.UUID     `json:"userUuid"`
	InvoiceUuid        uuid.UUID     `json:"invoiceUuid" gorm:"unique"`
	PaymentMethod      PaymentMethod `json:"paymentMethod"`
	AcquiringInvoiceId string        `json:"acquiringInvoiceId"`
	PaymentStatus      PaymentStatus `json:"paymentStatus"`
	CreatedAt          time.Time     `json:"createdAtt"`
	UpdatedAt          time.Time     `json:"updatedAt"`
}
