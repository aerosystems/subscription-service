package models

import (
	"github.com/google/uuid"
	"time"
)

type Invoice struct {
	Id                 int
	Amount             int
	UserUuid           uuid.UUID
	InvoiceUuid        uuid.UUID
	PaymentMethod      PaymentMethod
	AcquiringInvoiceId string
	PaymentStatus      PaymentStatus
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
