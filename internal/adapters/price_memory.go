package adapters

import (
	"errors"
	"github.com/aerosystems/subscription-service/internal/entities"
)

type PriceRepo struct {
	priceMap map[entities.SubscriptionType]map[entities.SubscriptionDuration]int
}

func NewPriceRepo() *PriceRepo {
	return &PriceRepo{
		priceMap: map[entities.SubscriptionType]map[entities.SubscriptionDuration]int{
			entities.StartupSubscriptionType: {
				entities.OneMonthSubscriptionDuration:    500,  // 5$ per month in cents
				entities.TwelveMonthSubscriptionDuration: 5000, // 50$ per year in cents
			},
			entities.BusinessSubscriptionType: {
				entities.OneMonthSubscriptionDuration:    1000,  // 10$ per month in cents
				entities.TwelveMonthSubscriptionDuration: 10000, // 100$ per year in cents
			},
		},
	}
}

func (pr *PriceRepo) GetPrice(subscriptionType entities.SubscriptionType, subscriptionDuration entities.SubscriptionDuration) (int, error) {
	price, ok := pr.priceMap[subscriptionType][subscriptionDuration]
	if !ok {
		return 0, errors.New("price not found")
	}
	return price, nil
}

func (pr *PriceRepo) GetAll() map[entities.SubscriptionType]map[entities.SubscriptionDuration]int {
	return pr.priceMap
}
