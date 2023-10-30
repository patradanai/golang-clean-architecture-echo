package main

import (
	"context"
	"fmt"
	"movie-service/configs"
	"movie-service/internal/errors"
	"movie-service/internal/handler"
	"movie-service/internal/repository"
	"movie-service/internal/usecase"
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
	mongoInstance := configs.InitMongo()

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
			e.Logger.Fatal("shutting down the server")
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
	// if c.Response().Committed {
	// 	return
	// }

	code := http.StatusInternalServerError
	var errHttp errs.Errors
	switch err.(type) {
	case *echo.HTTPError:
		// he, _ := err.(*echo.HTTPError)

		// return
		fmt.Println("1")
	case *errs.MetaError:
		he := err.(*errs.MetaError)

		errHttp = he

	default:
		errHttp = errs.WrapError(errors.InternalServerError, "")
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, errHttp.Error())
		}

		if err != nil {
			logger.Error(err)
		}
	}

}
