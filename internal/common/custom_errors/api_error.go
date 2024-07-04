package CustomErrors

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

type ApiError struct {
	Message  string
	HttpCode int
	GrpcCode codes.Code
}

func (e ApiError) Error() string {
	return e.Message
}

var (
	ErrInvalidCustomerUuid         = ApiError{Message: "Invalid customer uuid", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidSubscriptionType     = ApiError{Message: "Invalid subscription type", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidSubscriptionDuration = ApiError{Message: "Invalid subscription duration", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
)
