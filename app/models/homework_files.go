package models

import "time"

type HomeworkFiles struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	Type      int       `json:"type"`                               // 文件类型 1: 作业任务 2: 作业记录
	BindId    int       `json:"bindId"`                             // 绑定ID
	UserId    string    `json:"userId"`                             // 用户Id TeacherCard or StudentCard
	Filename  string    `json:"filename"`                           // 文件名
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`    // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`    // 更新时间
	Status    int       `json:"status"`                             // 状态
}

// HomeworkFilesReq 请求模型
type HomeworkFilesReq struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	Type     int    `json:"type"`                               // 文件类型 1: 作业任务 2: 作业记录
	BindId   int    `json:"bindId"`                             // 绑定ID
	UserId   string `json:"userId"`                             // 用户Id TeacherCard or StudentCard
	Filename string `json:"filename"`                           // 文件名
	Status   int    `json:"status"`                             // 状态
}

// HomeworkFilesResp 响应模型
type HomeworkFilesResp struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	Type      int       `json:"type"`                               // 文件类型 1: 作业任务 2: 作业记录
	BindId    int       `json:"bindId"`                             // 绑定ID
	UserId    string    `json:"userId"`                             // 用户Id TeacherCard or StudentCard
	Filename  string    `json:"filename"`                           // 文件名
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`    // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`    // 更新时间
	Status    int       `json:"status"`                             // 状态
}

// HomeworkFilesListResp 作业文件列表响应模型
type HomeworkFilesListResp struct {
	Records []HomeworkFilesResp `json:"records"` // 作业文件列表
	Total   int64               `json:"total"`   // 总记录数
}

func (HomeworkFilesResp) TableName() string {
	return "homework_files"
}
