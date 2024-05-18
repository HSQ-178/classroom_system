package models

import (
	"classroom-system/utils"
	"time"
)

type HomeworkRecords struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	HomeworkId  int       `json:"homeworkId"`                         // 作业编号
	StudentCard string    `json:"studentCard"`                        // 学生编号
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`    // 创建时间
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`    // 更新时间
	Status      int       `json:"status"`                             // 状态
}

// HomeworkRecordsReq 请求模型
type HomeworkRecordsReq struct {
	Id          int              `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	HomeworkId  int              `json:"homeworkId"`                         // 作业编号
	StudentCard string           `json:"studentCard"`                        // 学生编号
	Status      int              `json:"status"`                             // 状态
	Paginate    utils.Pagination `json:"pagination"`                         //分页
}

// HomeworkRecordsResp 响应模型
type HomeworkRecordsResp struct {
	Id          int                 `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	HomeworkId  int                 `json:"homeworkId"`                         // 作业编号
	StudentCard string              `json:"studentCard"`                        // 学生编号
	CreatedAt   time.Time           `json:"createdAt" gorm:"autoCreateTime"`    // 创建时间
	UpdatedAt   time.Time           `json:"updatedAt" gorm:"autoUpdateTime"`    // 更新时间
	Status      int                 `json:"status"`                             // 状态
	FileList    []HomeworkFilesResp `json:"fileList"`                           // 作业文件列表
}

// HomeworkRecordsListResp 作业记录列表响应模型
type HomeworkRecordsListResp struct {
	Records []HomeworkRecordsResp `json:"records"` // 作业记录列表
	Total   int64                 `json:"total"`   // 总记录数
}
