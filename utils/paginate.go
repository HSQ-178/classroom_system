package utils

import "gorm.io/gorm"

// Pagination 分页模型
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pagination.Page
		if page == 0 {
			page = 1
		}
		pageSize := pagination.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
