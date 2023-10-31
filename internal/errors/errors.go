package errors

import errs "movie-service/pkg/errors"

var (

	// Bussiness Error Dynamic Message Support Only %v
	ErrBadAuthorization = errs.MetaErrorBadRequest.NewError(1001, "permission not allowed")

	ErrUserNotFound = errs.MetaErrorBadRequest.NewError(1002, "user not found %v") // %s for Replace username
)
