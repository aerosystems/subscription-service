package HTTPServer

import (
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PriceResponse struct {
	Prices map[string]map[string]int `json:"prices" example:"{\"trial\":{\"1m\":0,\"12m\":0},\"startup\":{\"1m\":1000,\"12m\":10000},\"business\":{\"1m\":10000,\"12m\":100000}}"`
}

func ModelToPriceResponse(prices map[models.SubscriptionType]map[models.SubscriptionDuration]int) *PriceResponse {
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
// @Success 200 {object} map[string]map[string]int
// @Failure 500 {object} handlers.ErrorResponse
// @Router /v1/prices [get]
func (ph PaymentHandler) GetPrices(c echo.Context) error {
	return c.JSON(http.StatusOK, ModelToPriceResponse(ph.paymentUsecase.GetPrices()))
}
