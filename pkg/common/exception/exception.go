package exception

import (
	"usdw/internal/domain/constant"
)

type ErrorResponse struct {
	HTTPStatus *int        `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (err *ErrorResponse) Error() interface{} {
	return err.Message
}

var (
	DefaultErrorResponse = ErrorResponse{
		HTTPStatus: &constant.HTTPStatus500,
		Code:       constant.INTERNAL_ERROR,
		Message:    "Internal server error",
	}

	DefaultErrInternalServer = &ErrorResponse{
		HTTPStatus: &constant.HTTPStatus500,
		Code:       constant.INTERNAL_ERROR,
		Message:    "Internal server error",
	}

	DefaultErrBadRequest = &ErrorResponse{
		HTTPStatus: &constant.HTTPStatus400,
		Code:       constant.INVALID,
		Message:    "Bad request",
	}

	DefaultErrPermissionDenied = &ErrorResponse{
		HTTPStatus: &constant.HTTPStatus403,
		Code:       constant.PERMISSION_DENIED,
		Message:    "Permission denied",
	}

	DefaultErrNotFound = &ErrorResponse{
		HTTPStatus: &constant.HTTPStatus404,
		Code:       constant.NOT_FOUND,
		Message:    "Not found",
	}

	DefaultErrUnauthenticated = &ErrorResponse{
		HTTPStatus: &constant.HTTPStatus401,
		Code:       constant.UNAUTHENTICATED,
		Message:    "Unauthorized",
	}
)

func PanicLogging(err interface{}) {
	if err != nil {
		panic(err)
	}
}
