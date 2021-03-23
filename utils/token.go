package utils

import (
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-server/app/dto"
	"go-server/global"
	"time"
)

func GetToken() *jwt.Token {
	return jwt.New(jwt.SigningMethodHS256)
}

func SetToken(data dto.TokenModel) (string, error) {
	token := GetToken()
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = data.Username
	claims["uid"] = data.Uid
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(global.JwtTimeout)).Unix()
	return token.SignedString([]byte(global.JwtSecret))
}

func GetTokenModel(ctx *fiber.Ctx) (token dto.TokenModel, err error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims != nil {
		token.Username = claims["name"].(string)
		token.Uid = claims["uid"].(string)
		return token, nil
	} else {
		return dto.TokenModel{}, errors.New("Get  token model error.")
	}

}

func GetUid(ctx *fiber.Ctx) string {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["uid"].(string)
}
