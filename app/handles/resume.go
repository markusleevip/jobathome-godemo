package handles

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/app/models"
	"go-server/kit"
	"go-server/utils"
)

type Resume struct {
}

func (Resume) MyResume(ctx *fiber.Ctx) error {
	if token, err := utils.GetTokenModel(ctx); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg(err.Error()))
	} else {
		model := models.Resume{}
		model.Uid = token.Uid
		if table, err := model.GetResume(); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
		} else {
			fmt.Println(table)
			res := dto.ResumeRes{Uid: table.Uid, ResumeId: table.ResumeId, Content: table.Content, IsOpen: table.IsOpen}
			return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
		}
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
	fmt.Println(body)
	resume := models.Resume{Uid: uid, ResumeId: body.ResumeId, Content: body.Content, IsOpen: body.IsOpen}
	resume.Save()
	if table, err := resume.GetResume(); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("查询失败"))
	} else {
		res := dto.ResumeRes{ResumeId: table.ResumeId, Content: table.Content}
		return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
	}
}

func (Resume) ShowResume(ctx *fiber.Ctx) error {
	resumeId := ctx.Params("resumeId")
	model := models.Resume{}
	model.ResumeId = resumeId
	if table, err := model.GetOpenResume(); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
	} else {
		fmt.Println(table)
		res := dto.ResumeRes{Uid: table.Uid, ResumeId: table.ResumeId, Content: table.Content, IsOpen: table.IsOpen}
		return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
	}
}
