package HTTPServer

import (
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetSubscriptionResponse struct {
	Name       string `json:"name" example:"business"`
	Duration   string `json:"duration" example:"12m"`
	AccessTime string `json:"accessTime" example:"2021-09-01T00:00:00Z"`
}

func ModelToSubscriptionResponse(subscription *models.Subscription) *GetSubscriptionResponse {
	return &GetSubscriptionResponse{
		Name:       subscription.Type.String(),
		Duration:   subscription.Duration.String(),
		AccessTime: subscription.AccessTime.String(),
	}
}

// GetSubscriptions godoc
// @Summary Get subscriptions
// @Description get subscriptions by userUuid
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Security ServiceApiKeyAuth
// @Success 200 {object} GetSubscriptionResponse
// @Failure 401 {object} handlers.ErrorResponse
// @Failure 403 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /v1/subscriptions [get]
func (sh SubscriptionHandler) GetSubscriptions(c echo.Context) error {
	user, err := GetUserFromContext(c.Request().Context())
	if err != nil {
		return err
	}
	subscription, err := sh.subscriptionUsecase.GetSubscription(user.UUID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ModelToSubscriptionResponse(subscription))
}
