package api

import (
	errs "movie-service/pkg/errors"

	"github.com/go-resty/resty/v2"
)

func WrapRestError(res *resty.Response) errs.Errors {

	return nil
}
