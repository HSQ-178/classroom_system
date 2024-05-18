package models

import (
	"classroom-system/utils"
	"time"
)

type HomeworkTask struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	TeacherCard string    `json:"teacherCard"`                        // 老师编号
	CourseId    int       `json:"courseId"`                           // 课程编号
	Title       string    `json:"title"`                              // 作业标题
	Subscribe   string    `json:"subscribe"`                          // 作业内容
	StartedAt   time.Time `json:"startedAt"`                          // 开始时间
	EndedAt     time.Time `json:"endedAt"`                            // 结束时间
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`    // 创建时间
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`    // 更新时间
	Status      int       `json:"status"`                             // 状态
}

// HomeworkTaskReq 请求模型
type HomeworkTaskReq struct {
	Id          int              `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	TeacherCard string           `json:"teacherCard"`                        // 老师编号
	CourseId    int              `json:"courseId"`                           // 课程编号
	Title       string           `json:"title"`                              // 作业标题
	Subscribe   string           `json:"subscribe"`                          // 作业内容
	Status      int              `json:"status"`                             // 状态
	Paginate    utils.Pagination `json:"pagination"`                         //分页
}

// HomeworkTaskResp 响应模型
type HomeworkTaskResp struct {
	Id          int                 `json:"id" gorm:"primaryKey;autoIncrement"`                           // 主键ID
	TeacherCard string              `json:"teacherCard"`                                                  // 老师编号
	CourseId    int                 `json:"courseId"`                                                     // 课程编号
	Title       string              `json:"title"`                                                        // 作业标题
	Subscribe   string              `json:"subscribe"`                                                    // 作业内容
	StartedAt   time.Time           `json:"startedAt"`                                                    // 开始时间
	EndedAt     time.Time           `json:"endedAt"`                                                      // 结束时间
	CreatedAt   time.Time           `json:"createdAt" gorm:"autoCreateTime"`                              // 创建时间
	UpdatedAt   time.Time           `json:"updatedAt" gorm:"autoUpdateTime"`                              // 更新时间
	Status      int                 `json:"status"`                                                       // 状态
	FileList    []HomeworkFilesResp `json:"fileList" gorm:"references:Id;foreignKey:BindId"`              // 作业文件列表
	Teacher     TeacherResp         `json:"teacher" gorm:"references:TeacherCard;foreignKey:TeacherCard"` // 老师信息
	Course      CourseResp          `json:"course" gorm:"references:CourseId;foreignKey:Id"`              // 课程信息
}

// HomeworkTaskListResp 作业列表响应模型
type HomeworkTaskListResp struct {
	Records []HomeworkTaskResp `json:"records"` // 作业列表
	Total   int64              `json:"total"`   // 总记录数
}
