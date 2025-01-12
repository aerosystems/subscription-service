package GRPCServer

import (
	"context"
	"github.com/aerosystems/subscription-service/internal/common/protobuf/subscription"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SubscriptionHandler struct {
	subscriptionUsecase SubscriptionUsecase
	subscription.UnimplementedSubscriptionServiceServer
}

func NewSubscriptionHandler(subscriptionUsecase SubscriptionUsecase) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionUsecase: subscriptionUsecase,
	}
}

func (sh SubscriptionHandler) CreateFreeTrialSubscription(ctx context.Context, req *subscription.CreateFreeTrialSubscriptionRequest) (*subscription.CreateFreeTrialSubscriptionResponse, error) {
	sub, err := sh.subscriptionUsecase.CreateFreeTrial(req.CustomerUuid)
	if err != nil {
		return nil, err
	}
	return &subscription.CreateFreeTrialSubscriptionResponse{
		SubscriptionUuid: sub.Uuid.String(),
	}, nil
}

func (sh SubscriptionHandler) GetSubscription(ctx context.Context, req *subscription.GetSubscriptionRequest) (*subscription.GetSubscriptionResponse, error) {
	sub, err := sh.subscriptionUsecase.GetSubscription(uuid.MustParse(req.CustomerUuid))
	if err != nil {
		return nil, err
	}
	return &subscription.GetSubscriptionResponse{
		SubscriptionType: sub.Type.String(),
		AccessTime:       timestamppb.New(sub.AccessTime),
	}, nil
}

func (sh SubscriptionHandler) DeleteSubscription(ctx context.Context, req *subscription.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, sh.subscriptionUsecase.DeleteSubscription(uuid.MustParse(req.SubscriptionUuid))
}
