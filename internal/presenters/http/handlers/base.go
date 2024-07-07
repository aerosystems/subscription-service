package handlers

import (
	"github.com/go-playground/validator/v10"
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
