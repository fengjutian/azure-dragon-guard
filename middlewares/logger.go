package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	fmt.Printf("Request: %s %s\n", c.Method(), c.Path())
	return c.Next()
}
