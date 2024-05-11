package handlers

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/presenters/http/middleware"
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

type SubscriptionResponse struct {
	Name       string `json:"name" example:"business"`
	Duration   string `json:"duration" example:"12m"`
	AccessTime string `json:"accessTime" example:"2021-09-01T00:00:00Z"`
}

func ModelToSubscriptionResponse(subscription *models.Subscription) *SubscriptionResponse {
	return &SubscriptionResponse{
		Name:       subscription.Kind.String(),
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
func (sh SubscriptionHandler) GetSubscriptions(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*middleware.AccessTokenClaims)
	subscription, err := sh.subscriptionUsecase.GetSubscription(uuid.MustParse(accessTokenClaims.UserUuid))
	if err != nil {
		return sh.ErrorResponse(c, http.StatusInternalServerError, "could not find subscription", err)
	}
	return sh.SuccessResponse(c, http.StatusOK, "subscription successfully found", ModelToSubscriptionResponse(subscription))
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

type PriceResponse struct {
	Prices map[string]map[string]int `json:"prices" example:"{\"trial\":{\"1m\":0,\"12m\":0},\"startup\":{\"1m\":1000,\"12m\":10000},\"business\":{\"1m\":10000,\"12m\":100000}}"`
}

func ModelToPriceResponse(prices map[models.KindSubscription]map[models.DurationSubscription]int) *PriceResponse {
	priceResponse := make(map[string]map[string]int)
	for kind, durations := range prices {
		priceResponse[kind.String()] = make(map[string]int)
		for duration, price := range durations {
			priceResponse[kind.String()][duration.String()] = price
		}
	}
	return &PriceResponse{Prices: priceResponse}
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
	return sh.SuccessResponse(c, http.StatusOK, "prices for all available subscriptions, in cents", ModelToPriceResponse(sh.subscriptionUsecase.GetPrices()))
}
