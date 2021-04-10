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

// 好友列表
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

// 粉丝列表
// fans
func (Friend) MyFans(ctx *fiber.Ctx) error {
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
		if tables, err := model.GetFansPage(page); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
		} else {
			fmt.Println(tables)
			return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(tables))
		}
	}
}

// 添加关注
// addFollow
func (Friend) AddFollow(ctx *fiber.Ctx) error {
	if token, err := utils.GetTokenModel(ctx); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg(err.Error()))
	} else {
		friendReq := dto.Friend{}.FriendReq
		if err := ctx.BodyParser(&friendReq); err != nil {
			log.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		model := models.Friend{}
		model.Uid = token.Uid
		model.FUid = friendReq.FUid
		if err := model.Create(); err != nil {
			log.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(kit.Ok())
	}
}
