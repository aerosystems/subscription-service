package entities

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	Uuid         uuid.UUID
	CustomerUuid uuid.UUID
	Type         SubscriptionType
	Duration     SubscriptionDuration
	AccessTime   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SubscriptionType struct {
	slug        string
	accessCount int64
}

var (
	UnknownSubscriptionType  = SubscriptionType{"unknown", 0}
	TrialSubscriptionType    = SubscriptionType{"trial", 300}
	StartupSubscriptionType  = SubscriptionType{"startup", 1000}
	BusinessSubscriptionType = SubscriptionType{"business", -1}
)

func (k SubscriptionType) String() string {
	return k.slug
}

func (k SubscriptionType) GetAccessCount() int64 {
	return k.accessCount
}

func SubscriptionTypeFromString(kind string) SubscriptionType {
	switch kind {
	case TrialSubscriptionType.String():
		return TrialSubscriptionType
	case StartupSubscriptionType.String():
		return StartupSubscriptionType
	case BusinessSubscriptionType.String():
		return BusinessSubscriptionType
	default:
		return UnknownSubscriptionType
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

func SubscriptionDurationFromString(duration string) SubscriptionDuration {
	switch duration {
	case OneMonthSubscriptionDuration.String():
		return OneMonthSubscriptionDuration
	case TwelveMonthSubscriptionDuration.String():
		return TwelveMonthSubscriptionDuration
	default:
		return UnknownSubscriptionDuration
	}
}
