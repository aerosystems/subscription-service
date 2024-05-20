package usecases

import (
	"context"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
	"time"
)

type SubscriptionUsecase struct {
	subsRepo SubscriptionRepository
}

func NewSubscriptionUsecase(subsRepo SubscriptionRepository) *SubscriptionUsecase {
	return &SubscriptionUsecase{
		subsRepo: subsRepo,
	}
}

func NewSubscription(userUuid uuid.UUID, subscriptionType models.SubscriptionType, subscriptionDuration models.SubscriptionDuration) *models.Subscription {
	return &models.Subscription{
		UserUuid:   userUuid,
		Type:       subscriptionType,
		Duration:   subscriptionDuration,
		AccessTime: time.Now().Add(subscriptionDuration.GetTimeDuration()),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (ss SubscriptionUsecase) CreateSubscription(userUuid uuid.UUID, subscriptionType models.SubscriptionType, subscriptionDuration models.SubscriptionDuration) error {
	sub := NewSubscription(userUuid, subscriptionType, subscriptionDuration)
	ctx := context.Background()
	return ss.subsRepo.Create(ctx, sub)
}

func (ss SubscriptionUsecase) CreateFreeTrial(userUuid uuid.UUID, subscriptionType models.SubscriptionType) error {
	sub := NewSubscription(userUuid, subscriptionType, models.OneWeekSubscriptionDuration)
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
