package models

import "classroom-system/utils"

// Student 学生模型
type Student struct {
	Id          int64  `json:"id"`
	Grade       int64  `json:"grade"`
	College     string `json:"college"`
	Major       string `json:"major"`
	Class       string `json:"class"`
	StudentCard string `json:"studentCard"`
	Name        string `json:"name"`
}

// StudentReq 学生请求参数
type StudentReq struct {
	Id          int64  `json:"-"`           //主键Id
	TeacherCard string `json:"teacherCard"` //教师编号
	Grade       int64  `json:"grade"`       //年级
	College     string `json:"college"`     //学院
	Major       string `json:"major"`       //专业
	Class       string `json:"class"`       //班级
	StudentCard string `json:"studentCard"` //学号
	Name        string `json:"name"`        //姓名
	//Page        int    `json:"page"`        //当前页码
	//PageSize    int    `json:"pageSize"`    //总页数
	Pagination utils.Pagination `json:"pagination"`
	MajorId    int64            `json:"majorId"` //专业Id
}

// StudentResp  学生响应模型
type StudentResp struct {
	Grade       int64  `json:"grade"`
	College     string `json:"college"`
	Major       string `json:"major"`
	Class       string `json:"class"`
	StudentCard string `json:"studentCard"`
	Name        string `json:"name"`
}

// StudentListResp  学生列表响应模型
type StudentListResp struct {
	TotalCount int64         `json:"totalCount"`        //总记录数
	TotalPage  int           `json:"totalPage"`         //总页数
	Response   []StudentResp `json:"response" gorm:"-"` //学生列表
}
