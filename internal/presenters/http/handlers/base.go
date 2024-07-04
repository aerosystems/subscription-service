package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type BaseHandler struct {
	mode      string
	log       *logrus.Logger
	validator validator.Validate
}

func NewBaseHandler(
	log *logrus.Logger,
	mode string,
) *BaseHandler {
	return &BaseHandler{
		mode:      mode,
		log:       log,
		validator: validator.Validate{},
	}
}

// Response is the type used for sending JSON around
type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// SuccessResponse takes a response status code and arbitrary data and writes a json response to the client
func (h BaseHandler) SuccessResponse(c echo.Context, statusCode int, message string, data any) error {
	payload := Response{
		Message: message,
		Data:    data,
	}
	return c.JSON(statusCode, payload)
}
