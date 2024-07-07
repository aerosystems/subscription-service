package broker

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aerosystems/subscription-service/internal/models"
	PubSub "github.com/aerosystems/subscription-service/pkg/pubsub"
	"github.com/google/uuid"
	"time"
)

const (
	defaultTimeout = 2 * time.Second
)

type ProjectEventsAdapter struct {
	pubsubClient          *PubSub.Client
	topicId               string
	subName               string
	createProjectEndpoint string
	projectServiceApiKey  string
}

func NewProjectEventsAdapter(pubsubClient *PubSub.Client, topicId, subName, createProjectEndpoint, projectServiceApiKey string) *ProjectEventsAdapter {
	return &ProjectEventsAdapter{
		pubsubClient:          pubsubClient,
		topicId:               topicId,
		subName:               subName,
		createProjectEndpoint: createProjectEndpoint,
		projectServiceApiKey:  projectServiceApiKey,
	}
}

type CreateProjectEvent struct {
	CustomerUuid     string    `json:"customerUuid"`
	SubscriptionType string    `json:"subscriptionType"`
	AccessTime       time.Time `json:"accessTime"`
}

func (s ProjectEventsAdapter) PublishCreateProjectEvent(customerUuid uuid.UUID, subscriptionType models.SubscriptionType, accessTime time.Time) error {
	event := CreateProjectEvent{
		CustomerUuid:     customerUuid.String(),
		SubscriptionType: subscriptionType.String(),
		AccessTime:       accessTime,
	}
	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal create subscription event: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	topic := s.pubsubClient.Client.Topic(s.topicId)
	ok, err := topic.Exists(ctx)
	defer topic.Stop()
	if err != nil {
		return fmt.Errorf("failed to check if topic exists: %w", err)
	}
	if !ok {
		if _, err := s.pubsubClient.Client.CreateTopic(ctx, s.topicId); err != nil {
			return fmt.Errorf("failed to create topic: %w", err)
		}
	}

	sub := s.pubsubClient.Client.Subscription(s.subName)
	ok, err = sub.Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if subscription exists: %w", err)
	}
	if !ok {
		if _, err := s.pubsubClient.Client.CreateSubscription(ctx, s.subName, pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: 10 * time.Second,
			PushConfig: pubsub.PushConfig{
				Endpoint: s.createProjectEndpoint,
			},
		}); err != nil {
			return fmt.Errorf("failed to create subscription: %w", err)
		}
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: eventData,
	})

	if _, err := result.Get(ctx); err != nil {
		return fmt.Errorf("failed to publish create subscription event: %w", err)
	}

	return nil
}
