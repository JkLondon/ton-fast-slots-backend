package httpErrors

import (
	"Casino/config"
	"Casino/pkg/logger"
	"Casino/pkg/types/errorlist"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type responseMsg struct {
	Message string `json:"message"`
}

const (
	msgInternalServerError = "Internal server error"
)

func Init(config *config.Config, logger logger.Logger) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		statusCode := 0
		message := ""
		if errors.Is(err, errorlist.ErrWrongAuthData) {
			message = err.Error()
			statusCode = fiber.StatusUnauthorized
		} else if errors.Is(err, errorlist.ErrWrongAuthCredentials) {
			message = err.Error()
			statusCode = fiber.StatusUnauthorized
		} else if errors.Is(err, errorlist.ErrUnknownSessionType) {
			statusCode = fiber.StatusInternalServerError
		} else if errors.Is(err, errorlist.ErrNotFoundAuthData) {
			statusCode = fiber.StatusUnauthorized
			message = errorlist.ErrNotFoundAuthData.Error()
		} else if errors.Is(err, errorlist.ErrWrongPasswordUpdateData) {
			statusCode = fiber.StatusUnauthorized
			message = errorlist.ErrWrongPasswordUpdateData.Error()
		} else if errors.Is(err, errorlist.ErrServiceInCart) {
			statusCode = fiber.StatusConflict
			message = errorlist.ErrServiceInCart.Error()
		} else if errors.Is(err, errorlist.ErrThisUsernameAlreadyTaken) {
			statusCode = fiber.StatusConflict
			message = errorlist.ErrThisUsernameAlreadyTaken.Error()
		} else if errors.Is(err, errorlist.ErrNotFoundVerifyCode) {
			statusCode = fiber.StatusBadRequest
			message = errorlist.ErrNotFoundVerifyCode.Error()
		} else if errors.Is(err, errorlist.ErrWrongCurrentPassword) {
			statusCode = fiber.StatusBadRequest
			message = errorlist.ErrWrongCurrentPassword.Error()
		} else {
			statusCode = fiber.StatusInternalServerError
			if config.Server.ShowUnknownErrorsInResponse {
				message = fmt.Sprintf("на проде этого сообщения не будет: %s", err.Error())
			} else {
				logger.Error(fmt.Errorf("%s %v", c.OriginalURL(), err))
				message = msgInternalServerError
			}

		}

		return c.Status(statusCode).JSON(responseMsg{Message: message})
	}
}
