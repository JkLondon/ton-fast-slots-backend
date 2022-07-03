package server

import (
	gameHTTP "Casino/internal/game/delivery/http"
	gameUseCase "Casino/internal/game/usecase"
	apiMiddlewares "Casino/internal/middleware"
	userHTTP "Casino/internal/user/delivery/http"
	userRepository "Casino/internal/user/repository"
	userUseCase "Casino/internal/user/usecase"
	"Casino/pkg/tonapi"
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/patrickmn/go-cache"
)

func (s *Server) MapHandlers() error {
	userSessionCache := cache.New(cache.NoExpiration, cache.NoExpiration)
	tonSDK := tonapi.NewSDK(s.cfg.TON.BaseURL)

	userRepo := userRepository.NewUserRepository(s.pgDB, s.cfg, s.logger)

	userUC := userUseCase.NewUserUseCase(s.cfg, userRepo, userSessionCache, tonSDK, s.logger)
	gameUC := gameUseCase.NewGameUseCase(s.cfg, tonSDK, s.logger)

	err := userUC.SetBot()
	if err != nil {
		return err
	}
	userHandlers := userHTTP.NewUserHandlers(s.cfg, userUC, s.logger)
	gameHandlers := gameHTTP.NewGameHandlers(s.cfg, gameUC, s.logger)

	s.fiber.Use(
		requestid.New(),
		logger.New(),
		cors.New(cors.Config{
			AllowOrigins: s.cfg.Server.AllowOrigins,
			AllowHeaders: s.cfg.Server.AllowHeaders,
			AllowMethods: s.cfg.Server.AllowMethods,
		}),
	)

	mw := apiMiddlewares.NewMDWManager(s.cfg, userSessionCache, s.logger)
	userGroup := s.fiber.Group("user")
	gameGroup := s.fiber.Group("game")
	
	userHTTP.MapRoutes(userGroup, userHandlers, mw)
	gameHTTP.MapRoutes(gameGroup, gameHandlers, mw)

	return s.saveRoutesToFile()
}

func (s *Server) saveRoutesToFile() error {
	data, err := json.MarshalIndent(s.fiber.Stack(), "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("routes.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
