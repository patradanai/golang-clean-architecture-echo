package handler

import (
	"movie-service/internal/domain"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, hdls domain.IHandlers) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	e.GET("/user", hdls.UserHandler().GetUser)
}
