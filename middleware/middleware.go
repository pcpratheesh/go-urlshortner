package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var (
	HeaderApiKeyContext = "x-api-key"
)

// ApiKeyMiddleware
// An api-key validator middleware
func ApiKeyMiddleware(apiKey string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		key := ctx.Get(HeaderApiKeyContext)
		// check header api key is provided
		if key == "" {
			logrus.Error("missing header api key")
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "missing header api key",
			})
		}

		// check the provided api key and configured api are same
		if apiKey != key {
			logrus.Errorf("invalid header api key %v", key)
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid header api key",
			})
		}

		return ctx.Next()
	}
}
