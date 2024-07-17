package HttpServer

import (
	"github.com/aerosystems/subscription-service/internal/validators"
	"github.com/go-playground/validator/v10"
)

func (s *Server) setupValidator() {
	validator := validator.New()
	s.echo.Validator = &validators.CustomValidator{Validator: validator}
}
