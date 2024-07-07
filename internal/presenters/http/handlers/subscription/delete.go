package subscription

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// DeleteSubscription godoc
// @Summary Delete subscription
// @Description Delete subscription
// @Tags subscriptions
// @Accept  json
// @Produce application/json
// @Security BearerAuth
// @Param subscriptionId path string true "Subscription ID"
// @Failure 501 {object} echo.HTTPError
// @Router /v1/subscriptions/{subscriptionId} [delete]
func (sh Handler) DeleteSubscription(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "not implemented")
}
