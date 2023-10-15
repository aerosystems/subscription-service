package models

import (
	"errors"
	"time"
)

const (
	Startup  KindSubscription = "startup"
	Business KindSubscription = "business"
)

type Subscription struct {
	Id         uint             `json:"id"`
	Kind       KindSubscription `json:"kind"`
	UserId     uint             `json:"userId"`
	AccessTime uint             `json:"accessTime"`
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt"`
}

func NewSubscription(userId uint, kind string, accessTime uint) (*Subscription, error) {
	ks := KindSubscription(kind)
	if ks != Startup && ks != Business {
		return nil, errors.New("invalid kind of subscription")
	}
	return &Subscription{
		UserId:     userId,
		Kind:       ks,
		AccessTime: accessTime,
	}, nil
}

type KindSubscription string

type SubscriptionRepository interface {
	Create(subscription *Subscription) error
	GetByUserId(userId uint) (*Subscription, error)
	Update(subscription *Subscription) error
	Delete(id uint) error
}
