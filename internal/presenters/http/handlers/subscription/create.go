package subscription

import (
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
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

func ModelToCreateSubscriptionResponse(subscription *models.Subscription) *CreateSubscriptionResponseBody {
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
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 403 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /v1/subscriptions/create [post]
func (sh Handler) CreateSubscription(c echo.Context) error {
	var requestPayload CreateSubscriptionRequest
	if err := c.Bind(&requestPayload); err != nil {
		return CustomErrors.ErrInvalidRequestBody
	}
	subscription, err := sh.subscriptionUsecase.CreateSubscription(requestPayload.CustomerUuid, requestPayload.SubscriptionType, requestPayload.SubscriptionDuration)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, ModelToCreateSubscriptionResponse(subscription))
}
