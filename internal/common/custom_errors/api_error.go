package CustomErrors

import "net/http"

type ApiError struct {
	Message  string
	HttpCode int
}

func (e ApiError) Error() string {
	return e.Message
}

var (
	ErrInvalidCustomerUuid         = ApiError{Message: "invalid customer uuid", HttpCode: http.StatusBadRequest}
	ErrInvalidSubscriptionType     = ApiError{Message: "invalid subscription type", HttpCode: http.StatusBadRequest}
	ErrInvalidSubscriptionDuration = ApiError{Message: "invalid subscription duration", HttpCode: http.StatusBadRequest}
)
