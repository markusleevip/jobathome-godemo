package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-server/app/handles"
	"go-server/app/middleware"
)

func InitRoutes(app *fiber.App) {
	// 登录
	app.Get("/login", handles.Login)
	app.Post("/login", handles.Login)
	// 注册
	app.Post("/logon", handles.Logon)

	app.Post("/restricted", middleware.Auth(), handles.Restricted)

	app.Post("/checkToken", middleware.Auth(), handles.CheckToken)

}
