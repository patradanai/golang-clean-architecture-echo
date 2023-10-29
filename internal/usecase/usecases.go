package usecase

import (
	"movie-service/internal/repository"
	userService "movie-service/internal/usecase/user"
)

type (
	services struct {
		repositories repository.IRepositories
	}

	IServices interface {
		UserService() userService.IUsers
	}
)

func InitServices(repos repository.IRepositories) IServices {
	return &services{
		repositories: repos,
	}
}

func (s *services) UserService() userService.IUsers {
	return userService.Users{}
}
