package models

type DurationSubscription string

const (
	OneMonthDurationSubscription    DurationSubscription = "monthly"
	TwelveMonthDurationSubscription DurationSubscription = "annually"
)

func (d DurationSubscription) String() string {
	return string(d)
}
