package subscription

import "github.com/aerosystems/subs-service/internal/models"

const defaultTimeDuration = 60 * 60 * 24 * 14 // 14 days

type SubsService interface {
	CreateFreeTrial(userId uint, kind string) error
	GetSubscription(userId uint) (*models.Subscription, error)
}

type SubsServiceImpl struct {
	subsRepo models.SubscriptionRepository
}

func NewSubsService(subsRepo models.SubscriptionRepository) *SubsServiceImpl {
	return &SubsServiceImpl{
		subsRepo: subsRepo,
	}
}

func (ss *SubsServiceImpl) CreateFreeTrial(userId uint, kind string) error {
	sub, err := models.NewSubscription(userId, kind, defaultTimeDuration)
	if err != nil {
		return err
	}
	return ss.subsRepo.Create(sub)
}

func (ss *SubsServiceImpl) GetSubscription(userId uint) (*models.Subscription, error) {
	return ss.subsRepo.GetByUserId(userId)
}
