package HTTPServer

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// UpdateSubscription godoc
// @Summary Update subscription
// @Description Update subscription
// @Tags subscriptions
// @Accept  json
// @Produce application/json
// @Security BearerAuth
// @Param subscriptionId path string true "Subscription ID"
// @Failure 501 {object} handlers.ErrorResponse
// @Router /v1/subscriptions/{subscriptionId} [put]
func (h Handler) UpdateSubscription(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "not implemented")
}
