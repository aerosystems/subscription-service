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
	ErrInvalidRequestBody          = ApiError{Message: "Invalid request body", HttpCode: http.StatusUnprocessableEntity, GrpcCode: codes.InvalidArgument}
	ErrInvalidRequestPayload       = ApiError{Message: "Invalid request payload", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidCustomerUuid         = ApiError{Message: "Invalid customer uuid", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidSubscriptionType     = ApiError{Message: "Invalid subscription type", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidSubscriptionDuration = ApiError{Message: "Invalid subscription duration", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
)
