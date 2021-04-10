package models

import (
	"fmt"
	"go-server/app/dto"
	"go-server/global"
	"time"
)

type Friend struct {
	Uid       string    `json:"uid" gorm:"size:32;comment:用户id"`
	FUid      string    `json:"fUid" gorm:"size:32;comment:朋友id"`
	CreatedAt time.Time `json:"createdAt"`
}

type FriendDto struct {
	Friend
	NickName string `json:"nickName" gorm:"size:64;comment:用户昵称" `
}

func (Friend) TableName() string {
	return Prefix + "friend"
}

func (f *Friend) GetFollowPage(info dto.PageInfo) (result dto.PageResult, err error) {
	var list []FriendDto
	var count int64
	table := global.GDB.Table(f.TableName() + " f").Select([]string{"f.*", "a.nick_name"})
	table = table.Joins(fmt.Sprintf("LEFT JOIN %s a ON f.f_uid = a.uid ", Account{}.TableName()))
	table.Where("f.uid = ?", f.Uid).Count(&count)
	if err = table.Scopes(dto.Paginate(info)).Find(&list, "f.uid = ?", f.Uid).Error; err != nil {
		return result, err
	} else {
		return dto.PageResult{
			List:     list,
			Total:    count,
			Page:     info.Page,
			Pages:    dto.GetPages(count,info.PageSize),
			PageSize: info.PageSize,
		}, nil
	}
}

// 添加好友
func (f *Friend) Create() error {
	var count int64
	global.GDB.Table(f.TableName()).Select("uid").Where("uid = ? and f_uid = ?", f.Uid, f.FUid).Count(&count)
	if count > 0 {
		return ErrDataAlreadyExists
	}
	global.GDB.Create(f)
	return nil
}

// 粉丝列表
func (f *Friend) GetFansPage(info dto.PageInfo) (result dto.PageResult, err error) {
	var list []FriendDto
	var count int64
	table := global.GDB.Table(f.TableName() + " f").Select([]string{"f.*", "a.nick_name"})
	table = table.Joins(fmt.Sprintf("LEFT JOIN %s a ON f.uid = a.uid ", Account{}.TableName()))
	table.Where("f.f_uid = ?", f.Uid).Count(&count)
	if err = table.Scopes(dto.Paginate(info)).Find(&list, "f.f_uid = ?", f.Uid).Error; err != nil {
		return result, err
	} else {
		return dto.PageResult{
			List:     list,
			Total:    count,
			Page:     info.Page,
			Pages:    dto.GetPages(count,info.PageSize),
			PageSize: info.PageSize,
		}, nil
	}
}
