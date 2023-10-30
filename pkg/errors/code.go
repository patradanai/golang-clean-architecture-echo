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
