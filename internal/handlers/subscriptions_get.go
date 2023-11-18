package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) GetSubscriptions(c echo.Context) error {
	userUuid, err := uuid.Parse(c.Get("userUuid").(string))
	if err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, "invalid user uuid", err)
	}
	subscription, err := h.subscriptionService.GetSubscription(userUuid)
	if err != nil {
		return h.ErrorResponse(c, http.StatusInternalServerError, "could not found subscription", err)
	}
	return h.SuccessResponse(c, http.StatusOK, "subscription successfully found", subscription)
}
