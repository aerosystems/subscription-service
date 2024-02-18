package rest

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SubscriptionHandler struct {
	*BaseHandler
	subscriptionUsecase SubscriptionUsecase
}

func NewSubscriptionHandler(baseHandler *BaseHandler, subscriptionUsecase SubscriptionUsecase) *SubscriptionHandler {
	return &SubscriptionHandler{baseHandler, subscriptionUsecase}
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
func (sh SubscriptionHandler) GetSubscriptions(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*models.AccessTokenClaims)
	subscription, err := sh.subscriptionUsecase.GetSubscription(uuid.MustParse(accessTokenClaims.UserUuid))
	if err != nil {
		return sh.ErrorResponse(c, http.StatusInternalServerError, "could not find subscription", err)
	}
	return sh.SuccessResponse(c, http.StatusOK, "subscription successfully found", subscription)
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

// GetPrices godoc
// @Summary Get prices
// @Description get prices for all available subscriptions, in cents
// @Tags prices
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{data=map[string]map[string]int}
// @Failure 500 {object} Response
// @Router /v1/prices [get]
func (sh SubscriptionHandler) GetPrices(c echo.Context) error {
	return sh.SuccessResponse(c, http.StatusOK, "prices for all available subscriptions, in cents", sh.subscriptionUsecase.GetPrices())
}
