package domain

type IHandlers interface {
	UserHandler() IUserHandler
}
