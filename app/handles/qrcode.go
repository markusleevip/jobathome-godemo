package handles

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	qrcode "github.com/skip2/go-qrcode"
	"go-server/global"
	"go-server/kit"
)

type Qrcode struct {
}

// http://127.0.0.1:8001/qr?url=http://www.jobathome.cn
func QrTest(ctx *fiber.Ctx) error {
	url := ctx.FormValue("url", "http://www.jobathome.cn")
	fmt.Println(url)
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("生成Qrcode错误。"))
	}
	ctx.Response().Header.SetContentType("image/png")
	return ctx.Status(fiber.StatusOK).Send(png)
}

func GetResumeQr(ctx *fiber.Ctx) error {
	resumeId := ctx.Params(" ")
	fmt.Println("BaseUrl=" + global.BaseUrl)
	url := ctx.FormValue("url", global.BaseUrl)+"/show/resume/"+resumeId
	fmt.Println(url)
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 128)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("生成Qrcode错误。"))
	}
	ctx.Response().Header.SetContentType("image/png")
	return ctx.Status(fiber.StatusOK).Send(png)
}
