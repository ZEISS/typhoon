package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// DefaultErrorHandler is the default error handler.
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()

	// retrive the custom error code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{"error": message, "code": code})
}
