package services

import "github.com/aerosystems/subs-service/internal/models"

const defaultTimeDuration = 60 * 60 * 24 * 14 // 14 days

type SubsService interface {
	CreateFreeTrial(userId int, kind string) error
	GetSubscription(userId int) (*models.Subscription, error)
	DeleteSubscription(userId int) error
}

type SubsServiceImpl struct {
	subsRepo models.SubscriptionRepository
}

func NewSubsService(subsRepo models.SubscriptionRepository) *SubsServiceImpl {
	return &SubsServiceImpl{
		subsRepo: subsRepo,
	}
}

func (ss *SubsServiceImpl) CreateFreeTrial(userId int, kind string) error {
	sub, err := models.NewSubscription(userId, kind, defaultTimeDuration)
	if err != nil {
		return err
	}
	return ss.subsRepo.Create(sub)
}

func (ss *SubsServiceImpl) GetSubscription(userId int) (*models.Subscription, error) {
	return ss.subsRepo.GetByUserId(userId)
}

func (ss *SubsServiceImpl) DeleteSubscription(userId int) error {
	sub, err := ss.subsRepo.GetByUserId(userId)
	if err != nil {
		return err
	}
	return ss.subsRepo.Delete(sub.Id)

}
