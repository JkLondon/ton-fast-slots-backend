package middleware

import (
	"Casino/config"
	"Casino/pkg/logger"
	"Casino/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type MDWManager struct {
	cfg         *config.Config
	userSession *cache.Cache
	logger      logger.Logger
}

func NewMDWManager(
	cfg *config.Config,
	userSession *cache.Cache,
	logger logger.Logger,
) *MDWManager {
	return &MDWManager{cfg: cfg, userSession: userSession, logger: logger}
}

func (mw *MDWManager) AuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !mw.cfg.Middleware.Enable && c.Get("enableMDW") != "true" {
			c.Locals("TGID", mw.cfg.Middleware.TGID)
			return c.Next()
		}

		sessionData, found := mw.userSession.Get(utils.Clone(c.Get(fiber.HeaderAuthorization)))
		if !found {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		tgID, ok := sessionData.(int64)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Locals("tgID", tgID)
		return c.Next()
	}
}
