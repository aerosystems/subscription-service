package subscription

import (
	"encoding/json"
	"errors"
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
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 403 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /v1/subscriptions/create-free-trial [post]
func (sh Handler) CreateFreeTrial(c echo.Context) error {
	var requestPayload CreateFreeTrialRequest
	if err := c.Bind(&requestPayload); err != nil {
		return sh.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}

	var event CreateSubscriptionEvent
	if err := json.Unmarshal(requestPayload.Message.Data, &event); err != nil {
		return sh.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}

	subscription, err := sh.subscriptionUsecase.CreateFreeTrial(event.CustomerUuid)
	if err != nil {
		var apiErr CustomErrors.ApiError
		if errors.As(err, &apiErr) {
			return sh.ErrorResponse(c, apiErr.HttpCode, apiErr.Message, err)
		}
		return sh.ErrorResponse(c, http.StatusInternalServerError, "could not create user", err)
	}

	return c.JSON(http.StatusCreated, ModelToCreateSubscriptionResponse(subscription))
}
