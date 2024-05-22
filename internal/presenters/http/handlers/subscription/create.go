package subscription

import (
	"errors"
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateSubscriptionRequest struct {
	CreateSubscriptionRequestBody
}

type CreateSubscriptionRequestBody struct {
	CustomerUuid         string `json:"customerUuid"`
	SubscriptionType     string `json:"subscriptionType"`
	SubscriptionDuration string `json:"subscriptionDuration"`
}

type CreateSubscriptionResponseBody struct {
	CustomerUuid         string `json:"customerUuid"`
	SubscriptionType     string `json:"subscriptionType"`
	SubscriptionDuration string `json:"subscriptionDuration"`
	AccessTime           string `json:"accessTime"`
}

func ModelToCreateSubscriptionResponse(subscription *models.Subscription) *CreateSubscriptionResponseBody {
	return &CreateSubscriptionResponseBody{
		CustomerUuid:         subscription.UserUuid.String(),
		SubscriptionType:     subscription.Type.String(),
		SubscriptionDuration: subscription.Duration.String(),
		AccessTime:           subscription.AccessTime.String(),
	}
}

func (sh Handler) CreateSubscription(c echo.Context) error {
	var requestPayload CreateSubscriptionRequest
	if err := c.Bind(&requestPayload); err != nil {
		return sh.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}

	subscription, err := sh.subscriptionUsecase.CreateSubscription(requestPayload.CustomerUuid, requestPayload.SubscriptionType, requestPayload.SubscriptionDuration)
	if err != nil {
		var apiErr CustomErrors.ApiError
		if errors.As(err, &apiErr) {
			return sh.ErrorResponse(c, apiErr.HttpCode, apiErr.Message, err)
		}
		return sh.ErrorResponse(c, http.StatusInternalServerError, "could not create user", err)
	}

	return c.JSON(http.StatusCreated, ModelToCreateSubscriptionResponse(subscription))
}
