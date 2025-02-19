package entities

import (
	"github.com/aerosystems/common-service/customerrors"
	"google.golang.org/grpc/codes"
	"net/http"
)

var (
	ErrInvalidRequestBody          = customerrors.InternalError{Message: "Invalid request body", HttpCode: http.StatusUnprocessableEntity, GrpcCode: codes.InvalidArgument}
	ErrInvalidRequestPayload       = customerrors.InternalError{Message: "Invalid request payload", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidCustomerUuid         = customerrors.InternalError{Message: "Invalid customer uuid", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidSubscriptionType     = customerrors.InternalError{Message: "Invalid subscription type", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
	ErrInvalidSubscriptionDuration = customerrors.InternalError{Message: "Invalid subscription duration", HttpCode: http.StatusBadRequest, GrpcCode: codes.InvalidArgument}
)
