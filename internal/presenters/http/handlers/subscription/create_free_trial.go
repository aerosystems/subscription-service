package subscription

import (
	"encoding/json"
	CustomErrors "github.com/aerosystems/subscription-service/internal/common/custom_errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateFreeTrialRequest struct {
	CreateFreeTrialRequestBody
}

type CreateFreeTrialRequestBody struct {
	Message struct {
		Data []byte `json:"data"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type CreateSubscriptionEvent struct {
	CustomerUuid string `json:"customerUuid"`
}

// CreateFreeTrial godoc
// @Summary Create free trial
// @Description Create free trial
// @Tags subscriptions
// @Accept  json
// @Produce application/json
// @Security BearerAuth
// @Param raw body CreateFreeTrialRequestBody true "Create free trial"
// @Success 201 {object} CreateSubscriptionResponseBody
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 401 {object} handlers.ErrorResponse
// @Failure 403 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /v1/subscriptions/create-free-trial [post]
func (sh Handler) CreateFreeTrial(c echo.Context) error {
	var requestPayload CreateFreeTrialRequest
	if err := c.Bind(&requestPayload); err != nil {
		return CustomErrors.ErrInvalidRequestBody
	}
	var event CreateSubscriptionEvent
	if err := json.Unmarshal(requestPayload.Message.Data, &event); err != nil {
		return CustomErrors.ErrInvalidRequestPayload
	}
	subscription, err := sh.subscriptionUsecase.CreateFreeTrial(event.CustomerUuid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, ModelToCreateSubscriptionResponse(subscription))
}
