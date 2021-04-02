package models

import "time"

type Common struct {
	ID        uint `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
经历模板
*/
type Experience struct {
	Userid    string    `json:"userid" gorm:"size:32;comment:用户id"`
	StartTime time.Time `json:"startTime" gorm:"comment:开始时间"`
	EndTime   time.Time `json:"endTime" gorm:"comment:结束时间"`
	Sort      int       `json:"sort" gorm:"size:4;DEFAULT:0;"`
}

const (
	Prefix = "t_" // 表前缀
)

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
