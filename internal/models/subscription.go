package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	UserUuid   uuid.UUID            `json:"-"`
	Kind       KindSubscription     `json:"kind"`
	Duration   DurationSubscription `json:"duration"`
	AccessTime time.Time            `json:"accessTime"`
	CreatedAt  time.Time            `json:"createdAt"`
	UpdatedAt  time.Time            `json:"updatedAt"`
}

type KindSubscription struct {
	slug string
}

var (
	UnknownSubscription  = KindSubscription{"unknown"}
	TrialSubscription    = KindSubscription{"trial"}
	StartupSubscription  = KindSubscription{"startup"}
	BusinessSubscription = KindSubscription{"business"}
)

func (k KindSubscription) String() string {
	return k.slug
}

func NewKindSubscription(kind string) KindSubscription {
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

type DurationSubscription struct {
	slug string
}

var (
	UnknownDurationSubscription     = DurationSubscription{"unknown"}
	OneMonthDurationSubscription    = DurationSubscription{"1m"}
	TwelveMonthDurationSubscription = DurationSubscription{"12m"}
)

func (d DurationSubscription) String() string {
	return d.slug
}

func NewDurationSubscription(duration string) DurationSubscription {
	switch duration {
	case OneMonthDurationSubscription.String():
		return OneMonthDurationSubscription
	case TwelveMonthDurationSubscription.String():
		return TwelveMonthDurationSubscription
	default:
		return UnknownDurationSubscription
	}
}
