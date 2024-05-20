package subscription

import (
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/aerosystems/subscription-service/internal/presenters/http/middleware"
	"github.com/google/uuid"
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
// @Security ApiKeyAuth
// @Success 200 {object} Response{data=models.Subscription}
// @Failure 401 {object} Response
// @Failure 500 {object} Response
// @Router /v1/subscriptions [get]
func (sh Handler) GetSubscriptions(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*middleware.AccessTokenClaims)
	subscription, err := sh.subscriptionUsecase.GetSubscription(uuid.MustParse(accessTokenClaims.UserUuid))
	if err != nil {
		return sh.ErrorResponse(c, http.StatusInternalServerError, "could not find subscription", err)
	}
	return sh.SuccessResponse(c, http.StatusOK, "subscription successfully found", ModelToSubscriptionResponse(subscription))
}
