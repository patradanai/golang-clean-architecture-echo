package errs

import "net/http"

var (
	MetaErrorNotFound = ErrorDesc{
		httpCode: http.StatusNotFound,
	}

	MetaErrorUnauthorized = ErrorDesc{
		httpCode: http.StatusUnauthorized,
	}

	MetaErrorForbidden = ErrorDesc{
		httpCode: http.StatusForbidden,
	}

	// Internal Server Error
	MetaErrorInternalServer = ErrorDesc{
		httpCode: http.StatusInternalServerError,
	}

	// Biz Error
	MetaErrorBadRequest = ErrorDesc{
		httpCode: http.StatusBadRequest,
	}
)

var (
	ErrorCodeNotFound = MetaErrorInternalServer.NewError(-99999, "ErrorCode Not Found")

	InternalServerError = MetaErrorInternalServer.NewError(-1000, "Internal Server Error") // System Error
	BadRequest          = MetaErrorBadRequest.NewError(-1001, "Bad Request")

	ValidateFieldError = MetaErrorBadRequest.NewError(-1002, "Validate Field Errors")
)

type ErrorDesc struct {
	httpCode int
	code     int
	message  string
}

func (e *ErrorDesc) NewError(code int, message string) ErrorDesc {
	return ErrorDesc{
		code:    code,
		message: message,
	}
}
