package models

import (
	"classroom-system/utils"
)

// Student 学生模型
type Student struct {
	Id          int    `json:"id"`
	Grade       string `json:"grade"`
	College     string `json:"college"`
	Major       string `json:"major"`
	Class       string `json:"class"`
	StudentCard string `json:"studentCard"`
	Name        string `json:"name"`
	status      int    `json:"status"`
}

// StudentReq 学生请求参数
type StudentReq struct {
	Id          int              `json:"-"`           //主键Id
	TeacherCard string           `json:"teacherCard"` //教师编号
	Grade       string           `json:"grade"`       //年级
	College     string           `json:"college"`     //学院
	Major       string           `json:"major"`       //专业
	Class       string           `json:"class"`       //班级
	StudentCard string           `json:"studentCard"` //学号
	Name        string           `json:"name"`        //姓名
	Pagination  utils.Pagination `json:"pagination"`  //分页
	status      int              `json:"status"`      //状态  1:在读 2:休学 3:退学 4:毕业
}

// StudentResp  学生响应模型
type StudentResp struct {
	Id          int    `json:"id"`
	Grade       string `json:"grade"`
	College     string `json:"college"`
	Major       string `json:"major"`
	Class       string `json:"class"`
	StudentCard string `json:"studentCard"`
	Name        string `json:"name"`
	status      int    `json:"status"`
}

// StudentListResp  学生列表响应模型
type StudentListResp struct {
	TotalCount int64         `json:"totalCount"`        //总记录数
	TotalSize  int           `json:"totalSize"`         //每页数据
	Response   []StudentResp `json:"response" gorm:"-"` //学生列表
}
