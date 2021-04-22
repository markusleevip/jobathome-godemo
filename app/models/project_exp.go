package models

import (
	"go-server/app/dto"
	"go-server/global"
	"go-server/utils"
)

/**
项目经历
*/
type ProjectExp struct {
	Common
	Experience
	ProjectId   string `json:"projectId" gorm:"size:32;"`
	ProjectName string `json:"projectName" gorm:"size:64;comment:项目名称"`
	Content     string `json:"content" gorm:"type:varchar(4000);comment:正文"`
	IsOpen      bool   `json:"isOpen"  gorm:"size:1;DEFAULT:false;comment:是否公开"`
}

func (ProjectExp) TableName() string {
	return Prefix + "project_exp"
}

// 创建
func (t *ProjectExp) Save() {
	if t.ProjectId != "" {
		global.GDB.Where("project_id = ? ", t.ProjectId).Updates(&t)
	} else {
		t.ProjectId = utils.NewGenId()
		global.GDB.Create(t)
	}
}

func (t *ProjectExp) Delete() {
	global.GDB.Where("uid = ? and project_id = ? ",
		t.Uid, t.ProjectId).Delete(&t)
}

func (t *ProjectExp) List() (list []dto.ProjectExpRes) {
	global.GDB.Table(t.TableName()).Select("uid, project_id, project_name, content, is_open, "+
		" start_time, end_time").Where("uid = ? ", t.Uid).Order(" start_time desc ").Find(&list)
	return list
}
