package fire

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/aerosystems/subs-service/internal/models"
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
	UserUuid   string    `firestore:"user_uuid"`
	Kind       string    `firestore:"kind"`
	Duration   string    `firestore:"duration"`
	AccessTime time.Time `firestore:"access_time"`
	CreatedAt  time.Time `firestore:"created_at"`
	UpdatedAt  time.Time `firestore:"updated_at"`
}

func (s *SubscriptionFire) ToModel() *models.Subscription {
	return &models.Subscription{
		UserUuid:   uuid.MustParse(s.UserUuid),
		Kind:       models.NewKindSubscription(s.Kind),
		Duration:   models.NewDurationSubscription(s.Duration),
		AccessTime: s.AccessTime,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
	}
}

func ModelToSubscriptionFire(subscription *models.Subscription) *SubscriptionFire {
	return &SubscriptionFire{
		UserUuid:   subscription.UserUuid.String(),
		Kind:       subscription.Kind.String(),
		Duration:   subscription.Duration.String(),
		AccessTime: subscription.AccessTime,
		CreatedAt:  subscription.CreatedAt,
		UpdatedAt:  subscription.UpdatedAt,
	}
}

func (r *SubscriptionRepo) GetByUserUuid(ctx context.Context, userUuid uuid.UUID) (*models.Subscription, error) {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	var subscription SubscriptionFire
	iter := r.client.Collection("subscriptions").Where("user_uuid", "==", userUuid.String()).Documents(c)
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
	_, err := r.client.Collection("subscriptions").Doc(subscription.UserUuid.String()).Set(c, ModelToSubscriptionFire(subscription))
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, subscription *models.Subscription) error {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := r.client.Collection("subscriptions").Doc(subscription.UserUuid.String()).Set(c, ModelToSubscriptionFire(subscription))
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepo) Delete(ctx context.Context, subscription *models.Subscription) error {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := r.client.Collection("subscriptions").Doc(ModelToSubscriptionFire(subscription).UserUuid).Delete(c)
	if err != nil {
		return err
	}
	return nil
}
