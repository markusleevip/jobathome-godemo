package models

import (
	"go-server/global"
	"go-server/utils"
)

/**
	简历
 */
type Resume struct {
	Common
	Uid	 string `json:"uid" gorm:"size:32;comment:用户id"`
	ResumeId	string  `json:"resumeId" gorm:"size:32;"`
	Content		string `json:"content" gorm:"type:varchar(4000);comment:正文"`
	IsOpen		uint8  `json:"isOpen"  gorm:"size:1;DEFAULT:0;comment:是否公开"`
}

func (Resume) TableName() string {
	return Prefix+"resume"
}

func ( t *Resume ) GetResume() (table Resume , err error) {
	if err = global.GDB.First(&table, "uid = ?", t.Uid).Error; err != nil {
		return Resume{}, err
	} else {
		return table, nil
	}
}

func ( t *Resume) Save() {
	if t.ResumeId != "" {
		global.GDB.Where("resume_id = ?", t.ResumeId).Updates(t)
	}else {
		t.ResumeId = utils.NewGenId()
		global.GDB.Create(t)
	}

}