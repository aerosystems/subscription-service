package fire

import (
	"cloud.google.com/go/firestore"
	"context"
)

type SubscriptionRepo struct {
	client *firestore.Client
}

func NewSubscriptionRepo(client *firestore.Client) *SubscriptionRepo {
	return &SubscriptionRepo{
		client: client,
	}
}

func (r *SubscriptionRepo) GetByUserUuid(ctx context.Context, userUuid string) ([]map[string]interface{}, error) {
	var subscriptions []map[string]interface{}

	iter := r.client.Collection("subscriptions").Where("UserUuid", "==", userUuid).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		subscriptions = append(subscriptions, doc.Data())
	}

	return subscriptions, nil
}

func (r *SubscriptionRepo) GetByUuid(ctx context.Context, uuid string) (map[string]interface{}, error) {
	docRef := r.client.Collection("subscriptions").Doc(uuid)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	return doc.Data(), nil
}

func (r *SubscriptionRepo) Create(ctx context.Context, subscription map[string]interface{}) error {
	_, err := r.client.Collection("subscriptions").Doc(subscription["Uuid"].(string)).Set(ctx, subscription)
	return err
}

func (r *SubscriptionRepo) Update(ctx context.Context, subscription map[string]interface{}) error {
	_, err := r.client.Collection("subscriptions").Doc(subscription["Uuid"].(string)).Set(ctx, subscription)
	return err
}

func (r *SubscriptionRepo) Delete(ctx context.Context, subscription map[string]interface{}) error {
	_, err := r.client.Collection("subscriptions").Doc(subscription["Uuid"].(string)).Delete(ctx)
	return err
}
