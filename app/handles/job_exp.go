package handles

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/kit"
	"go-server/utils"
)

type JobExp struct {
}

func (JobExp) Save(ctx *fiber.Ctx) error {
	if token, err := utils.GetTokenModel(ctx); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg(err.Error()))
	} else {
		var body dto.JobExpRes
		ctx.BodyParser(&body)
		fmt.Println(token.Uid)
	}
	return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
}

func (JobExp) MyList(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
}
