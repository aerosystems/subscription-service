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

func NewSubscription(customerUuid uuid.UUID, subscriptionType models.SubscriptionType, subscriptionDuration models.SubscriptionDuration) *models.Subscription {
	return &models.Subscription{
		Uuid:         uuid.New(),
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
	if err := ss.subsRepo.Create(context.TODO(), sub); err != nil {
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
	if err := ss.subsRepo.Create(context.TODO(), sub); err != nil {
		return nil, fmt.Errorf("could not create subscription: %w", err)
	}
	return sub, nil
}

func (ss SubscriptionUsecase) GetSubscription(customerUuid uuid.UUID) (*models.Subscription, error) {
	return ss.subsRepo.GetByCustomerUuid(context.TODO(), customerUuid)
}

func (ss SubscriptionUsecase) DeleteSubscription(subscriptionUUID uuid.UUID) error {
	return ss.subsRepo.Delete(context.TODO(), subscriptionUUID)
}
