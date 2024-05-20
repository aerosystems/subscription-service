package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	UserUuid   uuid.UUID
	Type       SubscriptionType
	Duration   SubscriptionDuration
	AccessTime time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SubscriptionType struct {
	slug string
}

var (
	UnknownSubscription  = SubscriptionType{"unknown"}
	TrialSubscription    = SubscriptionType{"trial"}
	StartupSubscription  = SubscriptionType{"startup"}
	BusinessSubscription = SubscriptionType{"business"}
)

func (k SubscriptionType) String() string {
	return k.slug
}

func NewSubscriptionType(kind string) SubscriptionType {
	switch kind {
	case TrialSubscription.String():
		return TrialSubscription
	case StartupSubscription.String():
		return StartupSubscription
	case BusinessSubscription.String():
		return BusinessSubscription
	default:
		return UnknownSubscription
	}
}

type SubscriptionDuration struct {
	slug         string
	timeDuration time.Duration
}

var (
	UnknownSubscriptionDuration     = SubscriptionDuration{"unknown", 0}
	OneWeekSubscriptionDuration     = SubscriptionDuration{"1w", time.Hour * 24 * 7}
	OneMonthSubscriptionDuration    = SubscriptionDuration{"1m", time.Hour * 24 * 30}
	TwelveMonthSubscriptionDuration = SubscriptionDuration{"12m", time.Hour * 24 * 365}
)

func (d SubscriptionDuration) String() string {
	return d.slug
}

func (d SubscriptionDuration) GetTimeDuration() time.Duration {
	return d.timeDuration
}

func NewSubscriptionDuration(duration string) SubscriptionDuration {
	switch duration {
	case OneMonthSubscriptionDuration.String():
		return OneMonthSubscriptionDuration
	case TwelveMonthSubscriptionDuration.String():
		return TwelveMonthSubscriptionDuration
	default:
		return UnknownSubscriptionDuration
	}
}
