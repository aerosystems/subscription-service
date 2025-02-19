package adapters

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"time"
)

type SubscriptionRepo struct {
	client *firestore.Client
}

func NewSubscriptionRepo(client *firestore.Client) *SubscriptionRepo {
	return &SubscriptionRepo{
		client: client,
	}
}

type SubscriptionFire struct {
	Uuid         string    `firestore:"uuid"`
	CustomerUuid string    `firestore:"customer_uuid"`
	Type         string    `firestore:"type"`
	Duration     string    `firestore:"duration"`
	AccessTime   time.Time `firestore:"access_time"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}

func (s *SubscriptionFire) ToModel() *entities.Subscription {
	return &entities.Subscription{
		Uuid:         uuid.MustParse(s.Uuid),
		CustomerUuid: uuid.MustParse(s.CustomerUuid),
		Type:         entities.SubscriptionTypeFromString(s.Type),
		Duration:     entities.SubscriptionDurationFromString(s.Duration),
		AccessTime:   s.AccessTime,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

func ModelToSubscriptionFire(subscription *entities.Subscription) *SubscriptionFire {
	return &SubscriptionFire{
		Uuid:         subscription.Uuid.String(),
		CustomerUuid: subscription.CustomerUuid.String(),
		Type:         subscription.Type.String(),
		Duration:     subscription.Duration.String(),
		AccessTime:   subscription.AccessTime,
		CreatedAt:    subscription.CreatedAt,
		UpdatedAt:    subscription.UpdatedAt,
	}
}

func (r *SubscriptionRepo) GetByCustomerUuid(ctx context.Context, customerUuid uuid.UUID) (*entities.Subscription, error) {
	var subscription SubscriptionFire
	iter := r.client.Collection("subscriptions").Where("customer_uuid", "==", customerUuid.String()).Documents(ctx)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, errors.New("could not get subscription")
		}
		doc.DataTo(&subscription)
	}
	return subscription.ToModel(), nil
}

func (r *SubscriptionRepo) Create(ctx context.Context, subscription *entities.Subscription) error {
	_, err := r.client.Collection("subscriptions").Doc(subscription.Uuid.String()).Set(ctx, ModelToSubscriptionFire(subscription))
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, subscription *entities.Subscription) error {
	_, err := r.client.Collection("subscriptions").Doc(subscription.Uuid.String()).Set(ctx, ModelToSubscriptionFire(subscription))
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepo) Delete(ctx context.Context, subscriptionUUID uuid.UUID) error {
	_, err := r.client.Collection("subscriptions").Doc(subscriptionUUID.String()).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
