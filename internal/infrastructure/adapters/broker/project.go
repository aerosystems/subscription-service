package broker

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	PubSub "github.com/aerosystems/subscription-service/pkg/pubsub"
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
	CustomerUuid string `json:"customerUuid"`
}

func (s ProjectEventsAdapter) PublishCreateProjectEvent(customerUuid string) error {
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
			Topic: topic,
		}); err != nil {
			return fmt.Errorf("failed to create subscription: %w", err)
		}
	}

	event := CreateProjectEvent{
		CustomerUuid: customerUuid,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not marshal event: %w", err)
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Data: eventBytes,
	})
	if _, err := res.Get(ctx); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
