package usecases

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
	"time"
)

const defaultTimeDuration = 60 * 60 * 24 * 14 // 14 days in seconds

type SubscriptionUsecase struct {
	subsRepo  SubscriptionRepository
	priceRepo PriceRepository
}

func NewSubscriptionUsecase(subsRepo SubscriptionRepository, priceRepo PriceRepository) *SubscriptionUsecase {
	return &SubscriptionUsecase{
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

func (ss *SubscriptionUsecase) GetPrices() map[models.KindSubscription]map[models.DurationSubscription]int {
	return ss.priceRepo.GetAll()
}

func (ss *SubscriptionUsecase) CreateFreeTrial(userUuid uuid.UUID, kind models.KindSubscription) error {
	sub := NewSubscription(userUuid, kind, time.Now().Add(time.Second*defaultTimeDuration))
	return ss.subsRepo.Create(sub)
}

func (ss *SubscriptionUsecase) GetSubscription(userUuid uuid.UUID) (*models.Subscription, error) {
	return ss.subsRepo.GetByUserUuid(userUuid)
}

func (ss *SubscriptionUsecase) DeleteSubscription(userUuid uuid.UUID) error {
	sub, err := ss.subsRepo.GetByUserUuid(userUuid)
	if err != nil {
		return err
	}
	return ss.subsRepo.Delete(sub)
}
