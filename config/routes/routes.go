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

	app.Get("/qr", handles.QrTest)
	app.Get("/resumeQr/:resumeId", handles.GetResumeQr)

	app.Post("/restricted", middleware.Auth(), handles.Restricted)

	app.Post("/checkToken", middleware.Auth(), handles.CheckToken)

	my := app.Group("/my", middleware.Auth())
	my.Post("/resume", handles.Resume{}.MyResume)
	my.Put("/resume/save", handles.Resume{}.Save)
	show := app.Group("/show")
	show.Get("/resume/:resumeId", handles.Resume{}.ShowResume)

}
