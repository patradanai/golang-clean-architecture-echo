package handler

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, hdls IHandlers) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
}
