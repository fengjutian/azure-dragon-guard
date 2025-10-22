package handlers

import (
	"github.com/fengjutian/azure-dragon-guard/models"
	"github.com/fengjutian/azure-dragon-guard/services"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := services.GetAllUsers()
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	services.CreateUser(user)
	return c.JSON(user)
}
