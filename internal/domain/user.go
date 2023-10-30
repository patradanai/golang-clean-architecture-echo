package domain

import (
	errs "movie-service/pkg/errors"

	"github.com/labstack/echo/v4"
)

type IUserRepository interface {
	FindOne()
}

type IUserUsecase interface {
	FindById() errs.Errors
	FindOne()
	DeleteOne()
	UpdateOne()
}

type IUserHandler interface {
	GetUser(e echo.Context) error
}
