package dto

import (
	"gorm.io/gorm"
	"math"
)

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	Pages    int         `json:"pages" `
	PageSize int         `json:"pageSize"`
}

//分页封装
func Paginate(info PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if info.Page == 0 {
			info.Page = 1
		}
		switch {
		case info.PageSize > 1000:
			info.PageSize = 1000
		case info.PageSize <= 0:
			info.PageSize = 10
		}
		offset := (info.Page - 1) * info.PageSize
		return db.Offset(offset).Limit(info.PageSize)
	}
}

func GetPages(total int64, pageSize int) int {
	return int(math.Ceil(float64(total+int64(pageSize)-1)) / float64(pageSize))
}
