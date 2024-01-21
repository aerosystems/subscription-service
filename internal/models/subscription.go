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

type KindSubscription string

const (
	TrialSubscription    KindSubscription = "trial"
	StartupSubscription  KindSubscription = "startup"
	BusinessSubscription KindSubscription = "business"
)

func (k KindSubscription) String() string {
	return string(k)
}

type DurationSubscription string

const (
	OneMonthDurationSubscription    DurationSubscription = "1m"
	TwelveMonthDurationSubscription DurationSubscription = "12m"
)

func (d DurationSubscription) String() string {
	return string(d)
}
