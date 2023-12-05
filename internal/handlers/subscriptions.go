package handlers

import (
	"github.com/aerosystems/subs-service/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) GetSubscriptions(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*services.AccessTokenClaims)
	subscription, err := h.subscriptionService.GetSubscription(uuid.MustParse(accessTokenClaims.UserUuid))
	if err != nil {
		return h.ErrorResponse(c, http.StatusInternalServerError, "could not found subscription", err)
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
