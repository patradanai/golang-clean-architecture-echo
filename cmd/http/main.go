package main

import (
	"context"
	"fmt"
	"movie-service/internal/handler"
	"movie-service/internal/repository"
	"movie-service/internal/usecase"
	"movie-service/pkg/connector"
	cfg "movie-service/pkg/env"
	errs "movie-service/pkg/errors"
	"movie-service/pkg/logger"
	middlewares "movie-service/pkg/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func init() {
	cfg.InitEnv()
	logger.InitLogger()
}

func main() {
	logger.Info("Starting server...")

	e := echo.New()
	e.HTTPErrorHandler = handlerError

	// Middlewares
	e.Use(middlewares.InitMidleware()...)

	// Configs
	mongoInstance := connector.InitMongo(connector.Options{
		Uri: cfg.Get().Database.Uri,
	})

	// Repository
	repos := repository.InitRepositories(mongoInstance)

	// Service
	ucs := usecase.InitServices(repos)

	// Controller
	hdls := handler.InitHandlers(ucs)

	// Routes
	handler.InitRoutes(e, hdls)

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", cfg.Get().Http.Port)); err != http.ErrServerClosed {
			e.Logger.Fatalf("shutting down the server %v", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func handlerError(err error, c echo.Context) {
	var errHttp errs.Errors

	switch err.(type) {
	case *echo.HTTPError: // Echo Error
		he, _ := err.(*echo.HTTPError)
		errHttp = errs.WrapDError(he, errs.InternalServerError)
	case *errs.MetaError: // Error from MetaError
		he := err.(*errs.MetaError)
		errHttp = he
	default:
		errHttp = errs.WrapError(errs.ErrorCodeNotFound)
	}

	code := http.StatusInternalServerError
	if errHttp.HttpCode() != 0 {
		code = errHttp.HttpCode()
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, errs.ResponseError{
				Errors: errs.ErrorResponseBody{
					Code:     errHttp.Code(),
					Message:  errHttp.Message(),
					DescLine: errHttp.DescLine(),
					Fields:   errHttp.Stack(),
				},
			})
		}

		if err != nil {
			logger.Error(err)
		}
	}

}
