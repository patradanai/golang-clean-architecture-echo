package usecase

import (
	"movie-service/internal/domain"
	userService "movie-service/internal/usecase/user"
)

type (
	services struct {
		repositories domain.IRepositories
	}
)

func InitServices(repos domain.IRepositories) domain.IServices {
	return &services{
		repositories: repos,
	}
}

func (s *services) UserService() domain.IUserUsecase {
	return &userService.Users{}
}
