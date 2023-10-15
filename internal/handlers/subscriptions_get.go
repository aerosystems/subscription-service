package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) GetSubscriptions(c echo.Context) error {
	userId := c.Get("userId").(uint)
	subscription, err := h.SubscriptionService.GetSubscription(userId)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "could not found subscription", err)
	}
	return SuccessResponse(c, http.StatusOK, "subscription successfully found", subscription)
}
