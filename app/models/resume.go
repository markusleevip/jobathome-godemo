package models

import (
	"fmt"
	"go-server/app/dto"
	"go-server/global"
	"go-server/utils"
)

/**
简历
*/
type Resume struct {
	Common
	Uid      string `json:"uid" gorm:"size:32;comment:用户id"`
	ResumeId string `json:"resumeId" gorm:"size:32;"`
	Content  string `json:"content" gorm:"type:varchar(4000);comment:正文"`
	IsOpen   bool   `json:"isOpen"  gorm:"size:1;DEFAULT:false;comment:是否公开"`
	Visits   uint   `json:"visits" gorm:"size:32;default:0;comment:访问量"`
}

type ResumeDto struct {
	Content  string `json:"content"`
	NickName string `json:"nickName"`
	ResumeId string `json:"resumeId"`
	Uid      string `json:"uid"`
}

func (Resume) TableName() string {
	return Prefix + "resume"
}

func (t *Resume) GetResume() (table Resume, err error) {
	if err = global.GDB.First(&table, "uid = ?", t.Uid).Error; err != nil {
		return Resume{}, err
	} else {
		return table, nil
	}
}

func (t *Resume) Save() {
	if t.ResumeId != "" {
		global.GDB.Where("resume_id = ?", t.ResumeId).Updates(t)
		if !t.IsOpen {
			global.GDB.Model(Resume{}).Where("resume_id = ?", t.ResumeId).Update("is_open", false)
		}
	} else {
		t.ResumeId = utils.NewGenId()
		global.GDB.Create(t)
	}
}

func (t *Resume) GetOpenResume() (table Resume, err error) {
	if err = global.GDB.First(&table, "resume_id = ? and is_open = 1 ", t.ResumeId).Error; err != nil {
		return Resume{}, err
	} else {
		return table, nil
	}
}

// 更新访问量
func (t *Resume) UpdateVisits() {
	global.GDB.Model(t).Update("visits", t.Visits+1)
}

func (t *Resume) GetResumePage(info dto.PageInfo) (result dto.PageResult, err error) {
	var list []ResumeDto
	var count int64
	// select t.content,a.nick_name from t_resume t INNER JOIN t_account a ON t.uid = a.uid WHERE is_open = 1 ORDER BY t.updated_at DESC;
	table := global.GDB.Table(t.TableName() + " t ").Select([]string{"t.content", "a.nick_name", "t.resume_id", "a.uid"})
	table = table.Joins(fmt.Sprintf(" INNER JOIN %s a ON t.uid = a.uid ", Account{}.TableName()))
	table.Where(" is_open = 1 ").Count(&count)
	if err = table.Scopes(dto.Paginate(info)).Order(" t.updated_at DESC ").Find(&list, " is_open = 1 ").Error; err != nil {
		return result, err
	} else {
		return dto.PageResult{
			List:     list,
			Total:    count,
			Page:     info.PageSize,
			Pages:    dto.GetPages(count, info.PageSize),
			PageSize: info.PageSize,
		}, nil
	}

}
