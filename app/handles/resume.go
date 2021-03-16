package handles

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/app/models"
	"go-server/kit"
)

type Resume struct {
}

func (Resume) MyResume(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	uid := claims["uid"].(string)
	fmt.Println(string(ctx.Request().Header.RawHeaders()))
	fmt.Println(user)
	fmt.Println("uid=" + uid)
	fmt.Println("name=" + name)
	model := models.Resume{}
	model.Uid = uid
	if table, err := model.GetResume(); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
	} else {
		fmt.Println(table)
		res := dto.ResumeRes{Uid: table.Uid, ResumeId: table.ResumeId, Content: table.Content}
		return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
	}
}

func (Resume) Save(ctx *fiber.Ctx) error {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := claims["uid"].(string)

	var body dto.ResumeReq
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json.",
		})
	}
	resume := models.Resume{Uid: uid, ResumeId: body.ResumeId, Content: body.Content}
	resume.Save()
	if table, err := resume.GetResume(); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("查询失败"))
	}else {
		res := dto.ResumeRes{ResumeId: table.ResumeId, Content: table.Content}
		return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
	}
}
