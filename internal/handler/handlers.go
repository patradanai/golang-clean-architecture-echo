package handler

import (
	"movie-service/internal/domain"
	userhdl "movie-service/internal/handler/user"
)

type handlers struct {
	ucs domain.IServices
}

func InitHandlers(usecases domain.IServices) domain.IHandlers {
	return &handlers{ucs: usecases}
}

func (h *handlers) UserHandler() domain.IUserHandler {
	return &userhdl.UserHandler{
		Usecases: h.ucs,
	}
}
