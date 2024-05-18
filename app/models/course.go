package models

import "time"

// Course 课程模型
type Course struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"` //主键Id
	TeacherCard string    `json:"teacherCard"`                        //教师编号
	Grade       string    `json:"grade"`                              //年级
	College     string    `json:"college"`                            //学院
	Major       string    `json:"major"`                              //专业
	Class       string    `json:"class"`                              //班级
	Course      string    `json:"course"`                             //课程名
	StartTime   time.Time `json:"startTime"`                          //上课时间
	EndTime     time.Time `json:"endTime"`                            //下课时间
	Classroom   string    `json:"classroom"`                          //教室
}

// CourseReq 课程请求模型
type CourseReq struct {
	Id        int       `json:"id"`
	Grade     string    `json:"grade"`
	College   string    `json:"college"`
	Major     string    `json:"major"`
	Class     string    `json:"class"`
	Course    string    `json:"course"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Classroom string    `json:"classroom"`
	status    int       `json:"status"`

	TeacherCard string `json:"teacherCard"`
	StudentCard string `json:"studentCard"`
	StudentName string `json:"studentName"`
}

// CourseResp 课程响应模型
type CourseResp struct {
	Id        int       `json:"id"`
	Grade     string    `json:"grade"`
	College   string    `json:"college"`
	Major     string    `json:"major"`
	Class     string    `json:"class"`
	Course    string    `json:"course"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Classroom string    `json:"classroom"`

	TeacherCard string `json:"teacherCard"`
}

// CourseList 课程列表响应模型
type CourseListResp struct {
	TotalCount int64        `json:"totalCount"`          //课程总数
	CourseResp []CourseResp `json:"courseResp" gorm:"-"` //课程列表
}

func (CourseResp) TableName() string {
	return "courses"
}
