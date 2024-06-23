package usecases

import (
	"context"
	"fmt"
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
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

func (ss SubscriptionUsecase) CreateSubscription(userUuidStr, subscriptionTypeStr, subscriptionDurationStr string) (*models.Subscription, error) {
	userUuid, err := uuid.Parse(userUuidStr)
	if err != nil {
		return nil, CustomErrors.ErrInvalidCustomerUuid
	}
	subscriptionType := models.SubscriptionTypeFromString(subscriptionTypeStr)
	if subscriptionType == models.UnknownSubscriptionType {
		return nil, CustomErrors.ErrInvalidSubscriptionType
	}
	subscriptionDuration := models.SubscriptionDurationFromString(subscriptionDurationStr)
	if subscriptionDuration == models.UnknownSubscriptionDuration {
		return nil, CustomErrors.ErrInvalidSubscriptionDuration
	}
	sub := NewSubscription(userUuid, subscriptionType, subscriptionDuration)
	ctx := context.Background()
	if err := ss.subsRepo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("could not create subscription: %w", err)
	}
	return sub, nil
}

func (ss SubscriptionUsecase) CreateFreeTrial(userUuidStr string) (*models.Subscription, error) {
	userUuid, err := uuid.Parse(userUuidStr)
	if err != nil {
		return nil, CustomErrors.ErrInvalidCustomerUuid
	}
	sub := NewSubscription(userUuid, models.TrialSubscriptionType, models.OneWeekSubscriptionDuration)
	ctx := context.Background()
	if err := ss.subsRepo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("could not create subscription: %w", err)
	}
	return sub, nil
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
