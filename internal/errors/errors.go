package errors

import errs "movie-service/pkg/errors"

var (
	InternalServerError = errs.MetaErrorInternalServer.NewError(-1000, "Internal Server Error")
	BadRequest          = errs.MetaErrorBadRequest.NewError(-1001, "Bad Request")
)
