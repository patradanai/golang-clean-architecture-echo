package repository

import (
	"movie-service/internal/domain"
	userRepository "movie-service/internal/repository/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type repositories struct {
	*mongo.Client
}

func InitRepositories(mnInstance *mongo.Client) domain.IRepositories {
	return &repositories{mnInstance}
}

func (r *repositories) UserRepository() domain.IUserRepository {
	return &userRepository.UserRepository{}
}
