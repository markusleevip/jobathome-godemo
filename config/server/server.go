package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-server/config/routes"
	"go-server/initialize"
)

func InitServer() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05.999",
		TimeZone:   "Asia/Shanghai",
	}))

	routes.InitRoutes(app)
	app.Listen(initialize.Application.Host + ":" + initialize.Application.Port)
}
