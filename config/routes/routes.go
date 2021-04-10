package routes

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go-server/app/handles"
	"go-server/app/middleware"
	"go-server/global"
	"net/http"
)

var (
	static embed.FS
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
	my.Post("/follow", handles.Friend{}.MyFollow)
	my.Post("/fans", handles.Friend{}.MyFans)
	my.Put("/follow/add", handles.Friend{}.AddFollow)
	my.Put("/project/save", handles.ProjectExp{}.Save)
	my.Post("/project/list", handles.ProjectExp{}.MyList)

	app.Use("/res",filesystem.New(filesystem.Config{
		Root: http.Dir(global.ResPath),
	}))

	show := app.Group("/show")
	show.Get("/resume/:resumeId", handles.Resume{}.ShowResume)

}
