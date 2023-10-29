package middlewarelogger

import (
	"encoding/json"
	"movie-service/pkg/logger"

	"golang.org/x/exp/slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareLogger() echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			if slices.Contains([]string{"/ping"}, c.Request().URL.Path) {
				return true
			} else {
				return false
			}
		},
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			logger.Infof("[Request][Method: %s] [Ip Address: %s] [Path: %s] [Header: %s] [BODY: %s]", c.Request().Method, c.Request().RemoteAddr, c.Request().URL.Path, toString(c.Request().Header), string(reqBody))

			logger.Infof("[Response][Method: %s] [Ip Address: %s] [Path: %s] [Header: %s] [BODY: %s]", c.Request().Method, c.Request().RemoteAddr, c.Request().URL.Path, toString(c.Request().Header), string(resBody))
		},
	})
}

func toString(in any) (out string) {
	json, _ := json.Marshal(in)
	return string(json)
}
