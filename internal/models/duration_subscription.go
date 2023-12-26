package models

type DurationSubscription string

const (
	OneMonthDurationSubscription    DurationSubscription = "1m"
	TwelveMonthDurationSubscription DurationSubscription = "12m"
)

func (d DurationSubscription) String() string {
	return string(d)
}
