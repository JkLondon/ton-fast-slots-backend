package server

import (
	"Casino/config"
	"Casino/pkg/httpErrors"
	"Casino/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg    *config.Config
	fiber  *fiber.App
	pgDB   *sqlx.DB
	logger logger.Logger
}

func NewServer(
	cfg *config.Config,
	db *sqlx.DB,
	logger logger.Logger,
) *Server {
	return &Server{
		cfg:    cfg,
		fiber:  fiber.New(fiber.Config{ErrorHandler: httpErrors.Init(cfg, logger), DisableStartupMessage: true}),
		pgDB:   db,
		logger: logger,
	}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(); err != nil {
		s.logger.Fatalf("Cannot map handlers: %v", err)
	}

	go func() {
		s.logger.Infof("Starting HTTP server on port: %s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
		if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
			s.logger.Fatalf("Error starting HTTP server: ", err)
		}
	}()

	s.logger.Infof("Server is ready")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	err := s.fiber.Shutdown()
	if err != nil {
		s.logger.Error(err)
	} else {
		s.logger.Info("Fiber server exited properly")
	}

	return nil
}
