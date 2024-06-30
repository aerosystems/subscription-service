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
	subsRepo       SubscriptionRepository
	projectAdapter ProjectAdapter
}

func NewSubscriptionUsecase(subsRepo SubscriptionRepository, projectAdapter ProjectAdapter) *SubscriptionUsecase {
	return &SubscriptionUsecase{
		subsRepo:       subsRepo,
		projectAdapter: projectAdapter,
	}
}

func NewSubscription(customerUuid uuid.UUID, subscriptionType models.SubscriptionType, subscriptionDuration models.SubscriptionDuration) *models.Subscription {
	return &models.Subscription{
		CustomerUuid: customerUuid,
		Type:         subscriptionType,
		Duration:     subscriptionDuration,
		AccessTime:   time.Now().Add(subscriptionDuration.GetTimeDuration()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (ss SubscriptionUsecase) CreateSubscription(customerUuidStr, subscriptionTypeStr, subscriptionDurationStr string) (*models.Subscription, error) {
	customerUuid, err := uuid.Parse(customerUuidStr)
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
	sub := NewSubscription(customerUuid, subscriptionType, subscriptionDuration)
	ctx := context.Background()
	if err := ss.subsRepo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("could not create subscription: %w", err)
	}
	return sub, nil
}

func (ss SubscriptionUsecase) CreateFreeTrial(customerUuidStr string) (*models.Subscription, error) {
	customerUuid, err := uuid.Parse(customerUuidStr)
	if err != nil {
		return nil, CustomErrors.ErrInvalidCustomerUuid
	}
	sub := NewSubscription(customerUuid, models.TrialSubscriptionType, models.OneWeekSubscriptionDuration)
	if err := ss.projectAdapter.PublishCreateProjectEvent(customerUuidStr); err != nil {
		return nil, fmt.Errorf("could not publish create project event: %w", err)
	}
	ctx := context.Background()
	if err := ss.subsRepo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("could not create subscription: %w", err)
	}
	return sub, nil
}

func (ss SubscriptionUsecase) GetSubscription(customerUuid uuid.UUID) (*models.Subscription, error) {
	ctx := context.Background()
	return ss.subsRepo.GetByCustomerUuid(ctx, customerUuid)
}

func (ss SubscriptionUsecase) DeleteSubscription(customerUuid uuid.UUID) error {
	ctx := context.Background()
	sub, err := ss.subsRepo.GetByCustomerUuid(ctx, customerUuid)
	if err != nil {
		return err
	}
	ctx = context.Background()
	return ss.subsRepo.Delete(ctx, sub)
}
