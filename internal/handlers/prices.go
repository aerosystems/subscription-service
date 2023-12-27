package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetPrices godoc
// @Summary Get prices
// @Description get prices for all available subscriptions, in cents
// @Tags prices
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{data=map[string]map[string]int}
// @Failure 500 {object} Response
// @Router /v1/prices [get]
func (h *BaseHandler) GetPrices(c echo.Context) error {
	return h.SuccessResponse(c, http.StatusOK, "prices for all available subscriptions, in cents", h.subscriptionService.GetPrices())
}
