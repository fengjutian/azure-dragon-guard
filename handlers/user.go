package handlers

import (
	"github.com/fengjutian/azure-dragon-guard/services"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := services.GetAllUsers()
	// 返回带有状态信息的JSON响应
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "获取用户列表成功",
		"data":    users,
		"count":   len(users),
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(services.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	services.CreateUser(user)
	return c.JSON(user)
}
