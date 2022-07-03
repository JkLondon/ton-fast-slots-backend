package http

import (
	"Casino/config"
	"Casino/internal/user"
	"Casino/pkg/logger"
	"Casino/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type userHandlers struct {
	cfg    *config.Config
	userUC user.UseCase
	logger logger.Logger
}

func NewUserHandlers(
	cfg *config.Config,
	userUC user.UseCase,
	log logger.Logger,
) user.Handlers {
	return &userHandlers{cfg: cfg, userUC: userUC, logger: log}
}

func (u *userHandlers) GetSelfInfo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.StartFiberTrace(c, "userHandlers.GetSelfInfo")
		defer span.End()

		tgID, ok := c.Locals("TGID").(int64)
		if !ok {
			u.logger.Warn("???")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		result, err := u.userUC.GetSelfInfo(ctx, tgID)
		if err != nil {
			return err
		}

		return c.JSON(result)
	}
}
