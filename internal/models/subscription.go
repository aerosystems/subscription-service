package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	Id         int                  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserUuid   uuid.UUID            `json:"-" gorm:"unique"`
	Kind       KindSubscription     `json:"kind"`
	Duration   DurationSubscription `json:"duration"`
	AccessTime time.Time            `json:"accessTime"`
	CreatedAt  time.Time            `json:"createdAt"`
	UpdatedAt  time.Time            `json:"updatedAt"`
}
