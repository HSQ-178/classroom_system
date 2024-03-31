package models

import (
	"classroom-system/utils"
	"time"
)

// Record 记录模型
type Record struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	TeacherCard string    `json:"teacherCard"`
	CourseId    int       `json:"courseId"`
	StudentCard string    `json:"studentCard"`
	StudentName string    `json:"name"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	status      int       `json:"status"`
}

// RecordReq 记录请求参数
type RecordReq struct {
	Id          int    `json:"id"`          //主键ID
	TeacherCard string `json:"teacherCard"` //教师编号
	Grade       string `json:"grade"`       //年级
	College     string `json:"college"`     //学院
	Major       string `json:"major"`       //专业
	Class       string `json:"class"`       //班级
	CourseId    int    `json:"courseId"`    //课程ID
	StudentCard string `json:"studentCard"` //学生编号
	StudentName string `json:"name"`        //学生姓名
	CreateTime  string `json:"createTime"`  //创建时间
	Status      int    `json:"status"`      //状态  1：签到  2：缺勤  3：请假

	Paginate utils.Pagination `json:"pagination"` //分页
}

// RecordResp 记录响应模型
type RecordResp struct {
	Id          int       `json:"id"`          //主键ID
	TeacherCard string    `json:"teacherCard"` //教师编号
	Grade       string    `json:"grade"`       //年级
	College     string    `json:"college"`     //学院
	Major       string    `json:"major"`       //专业
	Class       string    `json:"class"`       //班级
	CourseId    int       `json:"courseId"`    //课程ID
	StudentCard string    `json:"studentCard"` //学生编号
	StudentName string    `json:"name"`        //学生姓名
	CreateTime  time.Time `json:"createTime"`  //创建时间
	Status      int       `json:"status"`      //状态  1：签到  2：缺勤  3：请假
}

// RecordRespList 记录签到或缺勤响应列表
type RecordSignInRespList struct {
	RecordResp []RecordResp `json:"recordResp"` //记录列表
	TotalCount int64        `json:"totalCount"` //总记录数
}

// NotSignInResp 未签到学生响应列表
type NotSignInResp struct {
	StudentList []Student `json:"studentList"` //学生列表
	TotalCount  int64     `json:"totalCount"`  //总记录数
}
