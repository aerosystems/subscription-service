package HTTPServer

import (
	"github.com/aerosystems/subs-service/internal/validators"
	"github.com/go-playground/validator/v10"
)

func (s *Server) setupValidator() {
	validator := validator.New()
	s.echo.Validator = &validators.CustomValidator{Validator: validator}
}
