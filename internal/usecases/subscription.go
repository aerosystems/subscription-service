package usecases

import (
	"context"
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

func (ss SubscriptionUsecase) GetPrices() map[models.KindSubscription]map[models.DurationSubscription]int {
	return ss.priceRepo.GetAll()
}

func (ss SubscriptionUsecase) CreateFreeTrial(userUuid uuid.UUID, kind models.KindSubscription) error {
	sub := NewSubscription(userUuid, kind, time.Now().Add(time.Second*defaultTimeDuration))
	ctx := context.Background()
	return ss.subsRepo.Create(ctx, sub)
}

func (ss SubscriptionUsecase) GetSubscription(userUuid uuid.UUID) (*models.Subscription, error) {
	ctx := context.Background()
	return ss.subsRepo.GetByUserUuid(ctx, userUuid)
}

func (ss SubscriptionUsecase) DeleteSubscription(userUuid uuid.UUID) error {
	ctx := context.Background()
	sub, err := ss.subsRepo.GetByUserUuid(ctx, userUuid)
	if err != nil {
		return err
	}
	ctx = context.Background()
	return ss.subsRepo.Delete(ctx, sub)
}
