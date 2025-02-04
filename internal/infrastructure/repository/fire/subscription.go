package fire

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"time"
)

const (
	defaultTimeout = 2 * time.Second
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

func (s *SubscriptionFire) ToModel() *models.Subscription {
	return &models.Subscription{
		Uuid:         uuid.MustParse(s.Uuid),
		CustomerUuid: uuid.MustParse(s.CustomerUuid),
		Type:         models.SubscriptionTypeFromString(s.Type),
		Duration:     models.SubscriptionDurationFromString(s.Duration),
		AccessTime:   s.AccessTime,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

func ModelToSubscriptionFire(subscription *models.Subscription) *SubscriptionFire {
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

func (r *SubscriptionRepo) GetByCustomerUuid(ctx context.Context, customerUuid uuid.UUID) (*models.Subscription, error) {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	var subscription SubscriptionFire
	iter := r.client.Collection("subscriptions").Where("customer_uuid", "==", customerUuid.String()).Documents(c)
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

func (r *SubscriptionRepo) Create(ctx context.Context, subscription *models.Subscription) error {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := r.client.Collection("subscriptions").Doc(subscription.Uuid.String()).Set(c, ModelToSubscriptionFire(subscription))
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, subscription *models.Subscription) error {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := r.client.Collection("subscriptions").Doc(subscription.Uuid.String()).Set(c, ModelToSubscriptionFire(subscription))
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepo) Delete(ctx context.Context, subscription *models.Subscription) error {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := r.client.Collection("subscriptions").Doc(ModelToSubscriptionFire(subscription).Uuid).Delete(c)
	if err != nil {
		return err
	}
	return nil
}
