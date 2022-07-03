package http

import (
	"Casino/config"
	"Casino/internal/game"
	"Casino/internal/models"
	"Casino/pkg/logger"
	"Casino/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type gameHandlers struct {
	cfg    *config.Config
	gameUC game.UseCase
	logger logger.Logger
}

func NewGameHandlers(
	cfg *config.Config,
	gameUC game.UseCase,
	log logger.Logger,
) game.Handlers {
	return &gameHandlers{cfg: cfg, gameUC: gameUC, logger: log}
}

func (g *gameHandlers) StartGame() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.StartFiberTrace(c, "userHandlers.StartGame")
		defer span.End()

		tgID, ok := c.Locals("TGID").(int64)
		if !ok {
			g.logger.Warn("???")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		result, err := g.gameUC.StartGame(ctx, tgID)
		if err != nil {
			return err
		}

		return c.JSON(result)
	}
}

func (g *gameHandlers) RoundSlot() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.StartFiberTrace(c, "userHandlers.StartGame")
		defer span.End()

		tgID, ok := c.Locals("TGID").(int64)
		if !ok {
			g.logger.Warn("???")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		params := models.RoundSlotParams{}
		if err := utils.ReadBodyRequest(c, &params); err != nil {
			return err
		}

		result, err := g.gameUC.RoundSlot(ctx, tgID, params)
		if err != nil {
			return err
		}

		return c.JSON(result)
	}
}

func (g *gameHandlers) EndGame() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.StartFiberTrace(c, "userHandlers.StartGame")
		defer span.End()

		tgID, ok := c.Locals("TGID").(int64)
		if !ok {
			g.logger.Warn("???")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		result, err := g.gameUC.EndGame(ctx, tgID)
		if err != nil {
			return err
		}

		return c.JSON(result)
	}
}
