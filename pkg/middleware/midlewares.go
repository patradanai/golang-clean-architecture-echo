package middlewares

import (
	middlewarelogger "movie-service/pkg/middleware/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMidleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.RequestID(),
		middleware.Recover(),
		middleware.CSRF(),
		middlewarelogger.MiddlewareLogger(),
	}
}

func CustomMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
