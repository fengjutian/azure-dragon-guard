package main

import (
	"github.com/fengjutian/azure-dragon-guard/database"
	"github.com/fengjutian/azure-dragon-guard/middlewares"
	"github.com/fengjutian/azure-dragon-guard/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 连接数据库并自动迁移
	database.ConnectDB()

	app := fiber.New()

	// 注册中间件
	app.Use(middlewares.Logger)

	// 注册路由
	routes.Setup(app)

	// 启动服务器
	app.Listen(":3000")
}
