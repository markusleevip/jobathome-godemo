package models

import (
	"go-server/app/dto"
	"go-server/global"
	"go-server/utils"
)

/**
工作经历
*/
type JobExp struct {
	Common
	Experience
	JobId       string `json:"jobId" gorm:"size:32;"`
	CompanyName string `json:"companyName" gorm:"size:64;comment:公司名称"`
	Content     string `json:"content" gorm:"type:varchar(4000);comment:正文"`
	IsOpen      bool   `json:"isOpen"  gorm:"size:1;DEFAULT:false;comment:是否公开"`
}

func (JobExp) TableName() string {
	return Prefix + "job_exp"
}

// 创建
func (t *JobExp) Save() {
	if t.JobId != "" {
		global.GDB.Where("job_id = ? ", t.JobId).Updates(&t)
	} else {
		t.JobId = utils.NewGenId()
		global.GDB.Create(t)
	}
}

func (t *JobExp) Delete() {
	global.GDB.Where("uid = ? and job_id = ? ",
		t.Uid, t.JobId).Delete(&t)
}

func (t *JobExp) List() (list []dto.JobExpRes) {
	global.GDB.Table(t.TableName()).Select("uid, job_id, company_name, content, is_open, "+
		" start_time, end_time").Where("uid = ? ", t.Uid).Order(" start_time desc ").Find(&list)
	return list
}
