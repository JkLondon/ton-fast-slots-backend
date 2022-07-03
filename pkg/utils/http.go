package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ReadBodyRequest(c *fiber.Ctx, out interface{}) error {
	if err := c.BodyParser(out); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), out)
}

func SendJsonBytes(c *fiber.Ctx, bytes []byte) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(bytes)
}

func SendJsonBytesIfExists(c *fiber.Ctx, bytes []byte) error {
	if len(bytes) == 0 {
		return c.SendStatus(fiber.StatusNoContent)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(bytes)
}
