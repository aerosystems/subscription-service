package HTTPServer

import (
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateSubscriptionRequest struct {
	CreateSubscriptionRequestBody
}

type CreateSubscriptionRequestBody struct {
	CustomerUuid         string `json:"customerUuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	SubscriptionType     string `json:"subscriptionType" example:"business"`
	SubscriptionDuration string `json:"subscriptionDuration" example:""`
}

type CreateSubscriptionResponseBody struct {
	CustomerUuid         string `json:"customerUuid"`
	SubscriptionType     string `json:"subscriptionType"`
	SubscriptionDuration string `json:"subscriptionDuration"`
	AccessTime           string `json:"accessTime"`
}

func ModelToCreateSubscriptionResponse(subscription *entities.Subscription) *CreateSubscriptionResponseBody {
	return &CreateSubscriptionResponseBody{
		CustomerUuid:         subscription.CustomerUuid.String(),
		SubscriptionType:     subscription.Type.String(),
		SubscriptionDuration: subscription.Duration.String(),
		AccessTime:           subscription.AccessTime.String(),
	}
}

// CreateSubscription godoc
// @Summary Create subscription
// @Description Create subscription
// @Tags subscriptions
// @Accept  json
// @Produce application/json
// @Security BearerAuth
// @Param raw body CreateSubscriptionRequestBody true "Create subscription"
// @Success 201 {object} CreateSubscriptionResponseBody
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 401 {object} handlers.ErrorResponse
// @Failure 403 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /v1/subscriptions/create [post]
func (h Handler) CreateSubscription(c echo.Context) error {
	var requestPayload CreateSubscriptionRequest
	if err := c.Bind(&requestPayload); err != nil {
		return entities.ErrInvalidRequestBody
	}
	subscription, err := h.subscriptionUsecase.CreateSubscription(c.Request().Context(), requestPayload.CustomerUuid, requestPayload.SubscriptionType, requestPayload.SubscriptionDuration)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, ModelToCreateSubscriptionResponse(subscription))
}
