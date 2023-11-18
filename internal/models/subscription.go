package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	Id         int              `json:"id" gorm:"primaryKey;autoIncrement"`
	UserUuid   uuid.UUID        `json:"userUuid" gorm:"unique"`
	Kind       KindSubscription `json:"kind"`
	AccessTime time.Time        `json:"accessTime"`
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt"`
}

func NewSubscription(userUuid uuid.UUID, kind KindSubscription, accessTime time.Time) (*Subscription, error) {
	return &Subscription{
		UserUuid:   userUuid,
		Kind:       kind,
		AccessTime: accessTime,
	}, nil
}

type SubscriptionRepository interface {
	Create(subscription *Subscription) error
	GetByUserUuid(userUuid uuid.UUID) (*Subscription, error)
	Update(subscription *Subscription) error
	Delete(subscription *Subscription) error
}
