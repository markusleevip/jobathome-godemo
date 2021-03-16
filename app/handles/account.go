package handles

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/app/models"
	"go-server/global"
	"go-server/kit"
	"go-server/utils"
	"time"

	"log"
)

type Account struct {
}

func Login(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Get("Content-Type"))
	fmt.Println(string(ctx.Body()))
	var body dto.LoginReq
	if err := ctx.BodyParser(&body); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json.",
		})
	}
	log.Println(body)
	log.Println("password=", body.Password)
	account := models.Account{Username: body.Username, Password: body.Password}
	if data, err := account.GetUser(); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("账号不存在。"))
	} else {
		if data.Password == utils.Sha1(body.Password+data.Salt) {
			// 账号密码验证成功
			token := utils.GetToken()
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = body.Username
			claims["uid"] = data.Uid
			claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
			if tokenKey, err := token.SignedString([]byte(global.JwtSecret)); err != nil {
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
	account := models.Account{Uid: utils.NewGenId(),Username: body.Username, Password: body.Password}
	if data, err := account.GetUser(); err != nil {
		// 进行注册操作
		log.Println(err)
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
