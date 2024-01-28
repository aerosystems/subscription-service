package rest

import (
	"github.com/aerosystems/subs-service/pkg/oauth_service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
func (h *BaseHandler) GetSubscriptions(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*OAuthService.AccessTokenClaims)
	subscription, err := h.subscriptionService.GetSubscription(uuid.MustParse(accessTokenClaims.UserUuid))
	if err != nil {
		return h.ErrorResponse(c, http.StatusInternalServerError, "could not find subscription", err)
	}
	return h.SuccessResponse(c, http.StatusOK, "subscription successfully found", subscription)
}

func (h *BaseHandler) CreateSubscription(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *BaseHandler) UpdateSubscription(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *BaseHandler) DeleteSubscription(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "not implemented")
}
