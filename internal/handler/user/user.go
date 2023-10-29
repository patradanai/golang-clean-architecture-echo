package userhdl

import "movie-service/internal/usecase"

type (
	UserHandler struct {
		Usecases usecase.IServices
	}

	IUserHandler interface {
		getUser()
	}
)

func (s *UserHandler) getUser() {

}
