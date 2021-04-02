package handles

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/app/models"
	"go-server/kit"
	"go-server/utils"
	"log"
	"strings"
)

type Account struct {
}

func Login(ctx *fiber.Ctx) error {
	var body dto.LoginReq
	if err := ctx.BodyParser(&body); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json.",
		})
	}

	userName := strings.ToLower(body.Username)
	account := models.Account{Username: userName, Password: body.Password}
	if data, err := account.GetUser(); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("账号不存在。"))
	} else {
		if data.Password == utils.Sha1(body.Password+data.Salt) {
			tokenModel := dto.TokenModel{Username: userName, Uid: data.Uid}
			if tokenKey, err := utils.SetToken(tokenModel); err != nil {
				log.Println(err)
				return ctx.Status(fiber.StatusBadRequest).JSON(
					kit.FailAndMsg(err.Error()))
			} else {
				res := dto.LoginRes{}
				res.Token = tokenKey
				res.Username = body.Username
				return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
			}
		} else {
			return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("账号或者密码错误。"))
		}
	}

	return nil

}

// 注册
func Logon(ctx *fiber.Ctx) error {
	var body dto.LogonReq
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json.",
		})
	}
	log.Println(body)
	res := dto.LogonRes{}
	res.Username = body.Username
	userName := strings.ToLower(body.Username)
	account := models.Account{Uid: utils.NewGenId(), Username: userName, Password: body.Password}
	if data, err := account.GetUser(); err != nil {
		// 进行注册操作
		log.Println(err)
		account.NickName = account.Username
		account.Create()
		return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
	} else {
		// 用户已经存在
		log.Println("data=", data)
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("账号已经存在。"))
	}

}

func CheckToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	fmt.Println(string(ctx.Request().Header.RawHeaders()))
	fmt.Println(user)
	res := dto.LoginRes{}
	res.Username = name
	return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(res))
}

func Restricted(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	fmt.Println(string(ctx.Request().Header.RawHeaders()))
	fmt.Println(user)
	return ctx.Status(fiber.StatusOK).JSON(kit.OkAndData(fmt.Sprintf("Hello %s Welcome to Job@Home.", name)))
}
