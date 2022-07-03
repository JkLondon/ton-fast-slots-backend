package game

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	StartGame() fiber.Handler
	RoundSlot() fiber.Handler
	EndGame() fiber.Handler
}
