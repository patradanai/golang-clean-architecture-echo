package userhdl

import (
	"movie-service/internal/domain"
	"movie-service/internal/errors"
	"movie-service/pkg/api"
	errs "movie-service/pkg/errors"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler struct {
		Usecases domain.IServices
	}

	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (s *UserHandler) GetUser(e echo.Context) error {
	return api.Response[User](e, "success", User{
		ID:   1,
		Name: "John Doe",
	})
}

func (s *UserHandler) FindUser(e echo.Context) error {
	return errs.WrapError(errors.ErrUserNotFound, "user not found")
}
