package repository

import (
	"errors"
	"github.com/aerosystems/subs-service/internal/models"
)

type PriceRepo struct {
	priceMap map[models.KindSubscription]map[models.DurationSubscription]int
}

func NewPriceRepo() *PriceRepo {
	return &PriceRepo{
		priceMap: map[models.KindSubscription]map[models.DurationSubscription]int{
			models.StartupSubscription: {
				models.OneMonthDurationSubscription:    500,  // 5$ per month in cents
				models.TwelveMonthDurationSubscription: 5000, // 50$ per year in cents
			},
			models.BusinessSubscription: {
				models.OneMonthDurationSubscription:    1000,  // 10$ per month in cents
				models.TwelveMonthDurationSubscription: 10000, // 100$ per year in cents
			},
		},
	}
}

func (pr *PriceRepo) GetPrice(kindSubscription models.KindSubscription, durationSubscription models.DurationSubscription) (int, error) {
	price, ok := pr.priceMap[kindSubscription][durationSubscription]
	if !ok {
		return 0, errors.New("price not found")
	}
	return price, nil
}

func (pr *PriceRepo) GetAll() map[models.KindSubscription]map[models.DurationSubscription]int {
	return pr.priceMap
}
