package models

import (
	"time"
)

type Authcode struct {
	id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	TeacherCard string    `json:"teacherCard"`
	Grade       string    `json:"grade"`      //年级
	College     string    `json:"college"`    //学院
	Major       string    `json:"major"`      //专业
	Class       string    `json:"class"`      //班级
	CourseId    int       `json:"courseId"`   //课程id
	Authcode    int       `json:"authcode"`   //签到码
	CreateTime  time.Time `json:"createTime"` //创建时间
}

// AuthcodeReq 请求模型
type AuthcodeReq struct {
	Id          int       `json:"id"`
	TeacherCard string    `json:"teacherCard"`
	Grade       string    `json:"grade"`
	College     string    `json:"college"`
	Major       string    `json:"major"`
	Class       string    `json:"class"`
	CourseId    int       `json:"courseId"`
	Authcode    int       `json:"authcode"`
	CreateTime  time.Time `json:"createTime"`
}

// 响应模型
type AuthcodeResp struct {
	Id          int       `json:"id"`
	TeacherCard string    `json:"teacherCard"`
	Grade       string    `json:"grade"`
	College     string    `json:"college"`
	Major       string    `json:"major"`
	Class       string    `json:"class"`
	CourseId    int       `json:"courseId"`
	Authcode    int       `json:"authcode"`
	CreateTime  time.Time `json:"createTime"`
}
