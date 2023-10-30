package domain

type IRepositories interface {
	UserRepository() IUserRepository
}
