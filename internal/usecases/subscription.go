package usecases

import (
	"context"
	"fmt"
	"github.com/aerosystems/subscription-service/internal/entities"
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

func NewSubscription(customerUuid uuid.UUID, subscriptionType entities.SubscriptionType, subscriptionDuration entities.SubscriptionDuration) *entities.Subscription {
	return &entities.Subscription{
		Uuid:         uuid.New(),
		CustomerUuid: customerUuid,
		Type:         subscriptionType,
		Duration:     subscriptionDuration,
		AccessTime:   time.Now().Add(subscriptionDuration.GetTimeDuration()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (ss SubscriptionUsecase) CreateSubscription(ctx context.Context, customerUUID, subscriptionTypeStr, subscriptionDurationStr string) (*entities.Subscription, error) {
	customerUuid, err := uuid.Parse(customerUUID)
	if err != nil {
		return nil, entities.ErrInvalidCustomerUuid
	}
	subscriptionType := entities.SubscriptionTypeFromString(subscriptionTypeStr)
	if subscriptionType == entities.UnknownSubscriptionType {
		return nil, entities.ErrInvalidSubscriptionType
	}
	subscriptionDuration := entities.SubscriptionDurationFromString(subscriptionDurationStr)
	if subscriptionDuration == entities.UnknownSubscriptionDuration {
		return nil, entities.ErrInvalidSubscriptionDuration
	}
	sub := NewSubscription(customerUuid, subscriptionType, subscriptionDuration)
	if err := ss.subsRepo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("could not create subscription: %w", err)
	}
	return sub, nil
}

func (ss SubscriptionUsecase) CreateFreeTrial(ctx context.Context, customerUUID string) (*entities.Subscription, error) {
	customerUuid, err := uuid.Parse(customerUUID)
	if err != nil {
		return nil, entities.ErrInvalidCustomerUuid
	}
	sub := NewSubscription(customerUuid, entities.TrialSubscriptionType, entities.OneWeekSubscriptionDuration)
	if err := ss.subsRepo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("could not create subscription: %w", err)
	}
	return sub, nil
}

func (ss SubscriptionUsecase) GetSubscription(ctx context.Context, customerUUID uuid.UUID) (*entities.Subscription, error) {
	return ss.subsRepo.GetByCustomerUuid(ctx, customerUUID)
}

func (ss SubscriptionUsecase) DeleteSubscription(ctx context.Context, subscriptionUUID uuid.UUID) error {
	return ss.subsRepo.Delete(ctx, subscriptionUUID)
}
