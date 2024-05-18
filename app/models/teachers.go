package models

// Teachers 老师模型
type Teachers struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	College     string `json:"college"`
	Major       string `json:"major"`
	TeacherCard string `form:"teacherCard" json:"teacherCard"`
	Name        string `json:"name"`
	Phone       string `json:"phone" gorm:"default: null"`
	Password    string `json:"password"`
}

type TeacherResp struct {
	Id          int    `json:"id"`
	College     string `json:"college"`
	Major       string `json:"major"`
	TeacherCard string `form:"teacherCard" json:"teacherCard"`
	Name        string `json:"name"`
	Phone       string `json:"phone" gorm:"default: null"`
}

// TeachersRegisterReq  老师注册请求参数
type TeachersRegisterReq struct {
	Id          int    `json:"-"`
	College     string `json:"college"`
	Major       string `json:"major"`
	TeacherCard string `json:"teacherCard"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

// TeachersLoginReq 老师登录请求参数
type TeachersLoginReq struct {
	TeacherCard string `json:"teacherCard"`
	Password    string `json:"password"`
}

// TeachersLoginResp 老师登录响应参数
type TeachersLoginResp struct {
	Token   string   `json:"token"`
	Teacher Teachers `json:"teacher"`
}

func (TeacherResp) TableName() string {
	return "teachers"
}
