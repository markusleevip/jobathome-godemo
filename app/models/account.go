package models

type Account struct {
	Common
	Username    string       `json:"userName" gorm:"size:64;comment:用户登录名"`
	Password    string       `json:"password"  gorm:"size:64;comment:用户登录密码"`
	Salt 		string 		 `json:"salt" gorm:"size:64;comment:加盐"`
	NickName    string       `json:"nickName" gorm:"size:64;comment:用户昵称" `
	Phone    string       	 `json:"phone" gorm:"size:20;comment:电话"`
	Email    string       	 `json:"email" gorm:"size:128;comment:邮箱"`
	Avatar   string       	  `json:"avatar" gorm:"size:128;default:x.jpg;comment:用户头像"`
	Status		int 		 `json:"status" gorm:"size:4;default:0;comment:用户状态"`
}