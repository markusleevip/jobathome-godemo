package handles

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/app/models"
	"go-server/kit"
	"go-server/utils"
	"log"
)

type Friend struct {
}

func (Friend) MyFollow(ctx *fiber.Ctx) error {
	if token, err := utils.GetTokenModel(ctx); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg(err.Error()))
	} else {
		page := dto.PageInfo{PageSize: 10}
		if err := ctx.BodyParser(&page); err != nil {
			log.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse json.",
			})
		}
		model := models.Friend{}
		model.Uid = token.Uid
		if tables, err := model.GetFollowPage(page); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
		} else {
			fmt.Println(tables)
			return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(tables))
		}
	}
}
