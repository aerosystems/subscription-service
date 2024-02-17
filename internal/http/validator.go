package HTTPServer

import (
	"github.com/aerosystems/auth-service/internal/validators"
	"github.com/go-playground/validator/v10"
)

func (s *Server) setupValidator() {
	validator := validator.New()
	validator.RegisterValidation("customPasswordRule", validators.CustomPasswordRule)
	s.echo.Validator = &validators.CustomValidator{Validator: validator}
}
