package adapters

import (
	"errors"
	"github.com/aerosystems/subscription-service/internal/models"
)

type PriceRepo struct {
	priceMap map[models.SubscriptionType]map[models.SubscriptionDuration]int
}

func NewPriceRepo() *PriceRepo {
	return &PriceRepo{
		priceMap: map[models.SubscriptionType]map[models.SubscriptionDuration]int{
			models.StartupSubscriptionType: {
				models.OneMonthSubscriptionDuration:    500,  // 5$ per month in cents
				models.TwelveMonthSubscriptionDuration: 5000, // 50$ per year in cents
			},
			models.BusinessSubscriptionType: {
				models.OneMonthSubscriptionDuration:    1000,  // 10$ per month in cents
				models.TwelveMonthSubscriptionDuration: 10000, // 100$ per year in cents
			},
		},
	}
}

func (pr *PriceRepo) GetPrice(subscriptionType models.SubscriptionType, subscriptionDuration models.SubscriptionDuration) (int, error) {
	price, ok := pr.priceMap[subscriptionType][subscriptionDuration]
	if !ok {
		return 0, errors.New("price not found")
	}
	return price, nil
}

func (pr *PriceRepo) GetAll() map[models.SubscriptionType]map[models.SubscriptionDuration]int {
	return pr.priceMap
}
