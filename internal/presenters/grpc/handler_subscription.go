package GRPCServer

import (
	"context"
	"github.com/aerosystems/subscription-service/internal/common/protobuf/subscription"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (sh SubscriptionHandler) CreateFreeTrialSubscription(context.Context, *subscription.CreateFreeTrialSubscriptionRequest) (*subscription.CreateFreeTrialSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFreeTrialSubscription not implemented")
}
func (sh SubscriptionHandler) DeleteSubscription(context.Context, *subscription.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
