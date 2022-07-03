package user

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	GetSelfInfo() fiber.Handler
}
