package models

import (
	"fmt"
	"go-server/global"
	"go-server/utils"
)

type Account struct {
	Common
	Uid      string `json:"uid" gorm:"size:32;comment:用户id"`
	Username string `json:"userName" gorm:"size:64;comment:用户登录名"`
	Password string `json:"password"  gorm:"size:64;comment:用户登录密码"`
	Salt     string `json:"salt" gorm:"size:64;comment:加盐"`
	NickName string `json:"nickName" gorm:"size:64;comment:用户昵称" `
	Phone    string `json:"phone" gorm:"size:20;comment:电话"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	Avatar   string `json:"avatar" gorm:"size:128;default:x.jpg;comment:用户头像"`
	Status   int    `json:"status" gorm:"size:4;default:0;comment:用户状态"`
	Age      int    `json:"age" gorm:"size:4;default:0;comment:年龄"`
}

func (Account) TableName() string {
	return Prefix + "account"
}

func (t *Account) Create() {
	t.Salt = utils.RandStr(16)
	fmt.Println(t.Salt)
	t.Password = utils.Sha1(t.Password + t.Salt)
	fmt.Println(t.Password)
	global.GDB.Create(t)
}

func (t *Account) GetUser() (data Account, err error) {
	if err = global.GDB.First(&data, "username = ?", t.Username).Error; err != nil {
		return Account{}, err
	} else {
		return data, nil
	}
}
