package http

import (
	"Casino/internal/game"
	"Casino/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapRoutes(group fiber.Router, h game.Handlers, mw *middleware.MDWManager) {
	group.Post("start_game", mw.AuthedMiddleware(), h.StartGame())
	group.Post("round_slot", mw.AuthedMiddleware(), h.RoundSlot())
	group.Post("end_game", mw.AuthedMiddleware(), h.EndGame())
}
