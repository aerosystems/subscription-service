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
