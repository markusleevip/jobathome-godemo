package handles

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/app/models"
	"go-server/kit"
	"go-server/utils"
	"time"
)

type ProjectExp struct {
}

func (ProjectExp) Save(ctx *fiber.Ctx) error {
	if token, err := utils.GetTokenModel(ctx); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg(err.Error()))
	} else {
		var body dto.ProjectExpReq
		ctx.BodyParser(&body)
		projectExp := models.ProjectExp{ProjectId: body.ProjectId, ProjectName: body.ProjectName,
			Content: body.Content, IsOpen: body.IsOpen}
		projectExp.StartTime, _ = time.Parse(ExpDateFormat, body.StartTime)
		projectExp.EndTime, _ = time.Parse(ExpDateFormat, body.EndTime)

		if !utils.CheckTime(projectExp.StartTime, projectExp.EndTime) {
			return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg(ErrExpDate.Error()))
		}
		projectExp.Uid = token.Uid
		projectExp.Save()
		fmt.Println(token.Uid)
	}
	return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
}

func (ProjectExp) MyList(ctx *fiber.Ctx) error {
	if token, err := utils.GetTokenModel(ctx); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg(err.Error()))
	} else {
		var list []dto.ProjectExpRes
		model := models.ProjectExp{}
		model.Uid = token.Uid
		list = model.List()
		fmt.Println(token.Uid)
		return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(list))
	}

}
