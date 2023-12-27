package services

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/google/uuid"
	"time"
)

const defaultTimeDuration = 60 * 60 * 24 * 14 // 14 days

type SubsServiceImpl struct {
	subsRepo  repository.SubscriptionRepository
	priceRepo repository.PriceRepository
}

func NewSubsServiceImpl(subsRepo repository.SubscriptionRepository, priceRepo repository.PriceRepository) *SubsServiceImpl {
	return &SubsServiceImpl{
		subsRepo:  subsRepo,
		priceRepo: priceRepo,
	}
}

func NewSubscription(userUuid uuid.UUID, kind models.KindSubscription, accessTime time.Time) *models.Subscription {
	return &models.Subscription{
		UserUuid:   userUuid,
		Kind:       kind,
		AccessTime: accessTime,
	}
}

func (ss *SubsServiceImpl) GetPrices() map[models.KindSubscription]map[models.DurationSubscription]int {
	return ss.priceRepo.GetAll()
}

func (ss *SubsServiceImpl) CreateFreeTrial(userUuid uuid.UUID, kind models.KindSubscription) error {
	sub := NewSubscription(userUuid, kind, time.Now().Add(time.Second*defaultTimeDuration))
	return ss.subsRepo.Create(sub)
}

func (ss *SubsServiceImpl) GetSubscription(userUuid uuid.UUID) (*models.Subscription, error) {
	return ss.subsRepo.GetByUserUuid(userUuid)
}

func (ss *SubsServiceImpl) DeleteSubscription(userUuid uuid.UUID) error {
	sub, err := ss.subsRepo.GetByUserUuid(userUuid)
	if err != nil {
		return err
	}
	return ss.subsRepo.Delete(sub)
}
