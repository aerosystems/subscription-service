package entities

import (
	"github.com/google/uuid"
	"time"
)

type Invoice struct {
	Id                 int
	Amount             int
	CustomerUuid       uuid.UUID
	InvoiceUuid        uuid.UUID
	PaymentMethod      PaymentMethod
	AcquiringInvoiceId string
	PaymentStatus      PaymentStatus
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
