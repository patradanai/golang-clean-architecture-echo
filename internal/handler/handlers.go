package handler

import (
	userhdl "movie-service/internal/handler/user"
	"movie-service/internal/usecase"
)

type handlers struct {
	ucs usecase.IServices
}

type IHandlers interface {
	UserHandler() userhdl.IUserHandler
}

func InitHandlers(usecases usecase.IServices) IHandlers {
	return &handlers{ucs: usecases}
}

func (h *handlers) UserHandler() userhdl.IUserHandler {
	return &userhdl.UserHandler{
		Usecases: h.ucs,
	}
}
