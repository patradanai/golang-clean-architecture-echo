package userhdl

import (
	"movie-service/internal/domain"
	errs "movie-service/pkg/errors"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler struct {
		Usecases domain.IServices
	}
)

func (s *UserHandler) GetUser(e echo.Context) error {
	return errs.WrapError(errs.MetaErrorInternalServer, "error")
}
