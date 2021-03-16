package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"go-server/global"
)

func Sha1(plain  string) string {
	secret := global.JwtSecret
	key := []byte(secret)
	pla := []byte(plain)
	mac := hmac.New(sha1.New, key)
	mac.Write(pla)
	cip := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(cip)
}

