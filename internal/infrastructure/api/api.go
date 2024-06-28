package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "tournament_participation_service/api"
	"tournament_participation_service/internal/application"
	"tournament_participation_service/internal/infrastructure/api/routes"

	echologrus "github.com/davrux/echo-logrus/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/echo-swagger"
)

// Server
// @title			Tournament Participation Service
// @version		1.0.0
// @description	Service for viewing participation in tournaments
//
// @host			127.0.0.1:30002
// @BasePath		/api/v1
type Server struct {
	srv *echo.Echo
}

func NewServer(app *application.App, logger *logrus.Logger) *Server {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	echologrus.Logger = logger
	e.Logger = echologrus.GetEchoLogger()
	e.Use(echologrus.Hook())

	e.GET("/docs/*", echoSwagger.WrapHandler)

	routes.Make(e.Group("/api/v1"), app)
	return &Server{srv: e}
}

func (s *Server) Start() {
	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		// Listen on application shutdown signals.
		s.srv.Logger.Info("Received a shutdown signal:", <-listener)

		// Shutdown HTTP server.
		if err := s.srv.Shutdown(context.Background()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.srv.Logger.Fatal(err)
		}
	}()

	// Start HTTP server.
	if err := s.srv.Start(":30002"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.srv.Logger.Fatal(err)
	}
}
