package subscription

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateSubscriptionRequest struct {
	CustomerUuid         string `json:"customerUuid"`
	SubscriptionType     string `json:"subscriptionType"`
	SubscriptionDuration string `json:"subscriptionDuration"`
}

func (sh Handler) CreateSubscription(c echo.Context) error {
	var requestPayload CreateSubscriptionRequest
	if err := c.Bind(&requestPayload); err != nil {
		return sh.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid request body", err)
	}

	return c.JSON(http.StatusNotImplemented, "not implemented")
}
