package GRPCServer

import (
	"context"
	"github.com/aerosystems/common-service/gen/protobuf/subscription"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SubscriptionService struct {
	subscriptionUsecase SubscriptionUsecase
	subscription.UnimplementedSubscriptionServiceServer
}

func NewSubscriptionService(subscriptionUsecase SubscriptionUsecase) *SubscriptionService {
	return &SubscriptionService{
		subscriptionUsecase: subscriptionUsecase,
	}
}

func (ss SubscriptionService) CreateFreeTrialSubscription(ctx context.Context, req *subscription.CreateFreeTrialSubscriptionRequest) (*subscription.CreateFreeTrialSubscriptionResponse, error) {
	sub, err := ss.subscriptionUsecase.CreateFreeTrial(ctx, req.CustomerUuid)
	if err != nil {
		return nil, err
	}
	return &subscription.CreateFreeTrialSubscriptionResponse{
		SubscriptionUuid: sub.Uuid.String(),
		SubscriptionType: sub.Type.String(),
		AccessTime:       timestamppb.New(sub.AccessTime),
		AccessCount:      sub.Type.GetAccessCount(),
	}, nil
}

func (ss SubscriptionService) GetSubscription(ctx context.Context, req *subscription.GetSubscriptionRequest) (*subscription.GetSubscriptionResponse, error) {
	sub, err := ss.subscriptionUsecase.GetSubscription(ctx, uuid.MustParse(req.CustomerUuid))
	if err != nil {
		return nil, err
	}
	return &subscription.GetSubscriptionResponse{
		SubscriptionType: sub.Type.String(),
		AccessTime:       timestamppb.New(sub.AccessTime),
		AccessCount:      sub.Type.GetAccessCount(),
	}, nil
}

func (ss SubscriptionService) DeleteSubscription(ctx context.Context, req *subscription.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, ss.subscriptionUsecase.DeleteSubscription(ctx, uuid.MustParse(req.SubscriptionUuid))
}
