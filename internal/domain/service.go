package domain

type IServices interface {
	UserService() IUserUsecase
}
