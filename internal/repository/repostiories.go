package repository

import (
	userRepository "movie-service/internal/repository/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type repositories struct {
	*mongo.Client
}

type IRepositories interface {
	UserRepository() userRepository.IUserRepository
}

func InitRepositories(mnInstance *mongo.Client) IRepositories {
	return &repositories{mnInstance}
}

func (r *repositories) UserRepository() userRepository.IUserRepository {
	return &userRepository.UserRepository{}
}
