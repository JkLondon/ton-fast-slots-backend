package http

import (
	"Casino/internal/middleware"
	"Casino/internal/user"

	"github.com/gofiber/fiber/v2"
)

func MapRoutes(group fiber.Router, h user.Handlers, mw *middleware.MDWManager) {
	group.Get("get_self_info", mw.AuthedMiddleware(), h.GetSelfInfo())
}
